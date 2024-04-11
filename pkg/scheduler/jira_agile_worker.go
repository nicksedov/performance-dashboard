package scheduler

import (
	"encoding/json"
	"fmt"
	"log"
	"performance-dashboard/pkg/profiles"
	database "performance-dashboard/pkg/database"
	dbmodel "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/http"
	jiramodel "performance-dashboard/pkg/jira/model"

	"gorm.io/gorm"
)

const (
	storyPointCustomType string = "com.pyxis.greenhopper.jira:jsw-story-points"
	dateTimeCustomType   string = "com.atlassian.jira.plugin.system.customfieldtypes:datetime"
	datePickerCustomType string = "com.atlassian.jira.plugin.system.customfieldtypes:datepicker"
	strintCustomType     string = "com.pyxis.greenhopper.jira:gh-sprint"
)

func jiraAgileWorker() error {

	config := profiles.GetSettings()
	boardId := config.JiraConfig.BoardID

	// Get active sprint
	getSprintApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint?state=active", boardId)
	sprint := jira.QueryPaged("GET", getSprintApiPath, &jiramodel.Sprint{})

	log.Printf("Active Sprint: %d\n", sprint.ID)

	database.SaveSprint(sprint)

	getIssuesApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint/%d/issue", boardId, sprint.ID)
	issues := jira.QueryOne("GET", getIssuesApiPath, &jiramodel.Issues{})

	issueCustomFields, _ := database.Read[dbmodel.IssueMetadata](
		func(items *[]dbmodel.IssueMetadata, db *gorm.DB) {
			db.Where("key like 'customfield_%'").Find(items)
		})

	for _, issue := range issues.Issues {
		fields := jiramodel.IssueFields{}
		fieldsJson, _ := json.Marshal(issue.Fields)
		json.Unmarshal(fieldsJson, &fields)
		// Process custom fields
		for _, customField := range *issueCustomFields {
			fieldVal := issue.Fields[customField.Key]
			if fieldVal == nil {
				continue
			}
			if customField.Custom == storyPointCustomType {
				fields.StoryPoints = fieldVal.(float64)
			} else if customField.Custom == dateTimeCustomType {
				if customField.Name == "Actual start" {
					fields.ActualStart = fieldVal.(string)
				} else if customField.Name == "Actual end" {
					fields.ActualEnd = fieldVal.(string)
				}
			} else if customField.Custom == datePickerCustomType {
				if customField.Name == "Start date" {
					fields.StartDate = fieldVal.(string)
				}
			} else if customField.Custom == strintCustomType {
				if customField.Name == "Sprint" {
					sprints := []jiramodel.IssueSprint{}
					sprintsJson, _ := json.Marshal(fieldVal)
					json.Unmarshal(sprintsJson, &sprints)
					fields.Sprints = sprints
				}
			}
		}
		database.SaveIssue(&issue, &fields)

	}
	return nil
}
