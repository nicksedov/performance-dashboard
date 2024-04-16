package scheduler

import (
	"encoding/json"
	"fmt"
	database "performance-dashboard/pkg/database"
	dbmodel "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/http"
	jiramodel "performance-dashboard/pkg/jira/model"
	"performance-dashboard/pkg/profiles"
	"time"

	"gorm.io/gorm"
)

const (
	storyPointCustomType string = "com.pyxis.greenhopper.jira:jsw-story-points"
	dateTimeCustomType   string = "com.atlassian.jira.plugin.system.customfieldtypes:datetime"
	datePickerCustomType string = "com.atlassian.jira.plugin.system.customfieldtypes:datepicker"
	sprintCustomType     string = "com.pyxis.greenhopper.jira:gh-sprint"
)

func jiraAgileWorker() error {

	config := profiles.GetSettings()
	boardId := config.JiraConfig.BoardID

	// Get active sprint
	getSprintApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint", boardId)
	sprints := jira.QueryPaged("GET", getSprintApiPath, &[]jiramodel.Sprint{})

	var activeSprintId int
	for _, sprint := range *sprints {
		if sprint.State == "active" {
			activeSprintId = sprint.ID
		}
		database.SaveSprint(&sprint)
	}

	poll, _ := database.NewPoll(activeSprintId)

	getIssuesApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint/%d/issue", boardId, activeSprintId)
	issues := jira.QueryOne("GET", getIssuesApiPath, &jiramodel.Issues{})
	customFieldsByIssueType := getCustomFields()

	for _, issue := range issues.Issues {
		issueId, subtasks := deepSaveIssue(poll, &issue, customFieldsByIssueType, 0)
		for _, subtask := range subtasks {
			time.Sleep(200 * time.Millisecond)
			subtaskDetails := jira.QueryOne("GET", subtask.Self, &jiramodel.Issue{})
			deepSaveIssue(poll, subtaskDetails, customFieldsByIssueType, issueId)
		}
	}
	return nil
}

func getCustomFields() *map[string][]dbmodel.IssueMetadata {
	customFields, _ := database.Read[dbmodel.IssueMetadata](
		func(items *[]dbmodel.IssueMetadata, db *gorm.DB) {
			db.Where("key like 'customfield_%'").Find(items)
		})
	customFieldsByIssueType := make(map[string][]dbmodel.IssueMetadata)
	for _, customField := range *customFields {
		issueTypeName := customField.IssueTypeName
		untranslatedName := customField.UntranslatedName
		keys := make([]string, 0, 2) 
		keys = append(keys, issueTypeName) 
		if untranslatedName != "" && untranslatedName != issueTypeName {
			keys = append(keys, untranslatedName)
		}
		for _, key := range keys {
			if customFieldsByIssueType[key] == nil {
				customFieldsByIssueType[key] = make([]dbmodel.IssueMetadata, 0, len(*customFields))
			} 
			cf := customFieldsByIssueType[key]
			customFieldsByIssueType[key] = append(cf, customField)
		}
	}
	return &customFieldsByIssueType
}

func deepSaveIssue(poll *dbmodel.Poll, issue *jiramodel.Issue, customFieldsByIssueType *map[string][]dbmodel.IssueMetadata, parentId int) (int, []jiramodel.Issue) {
	fields := jiramodel.IssueFields{}
	fieldsJson, _ := json.Marshal(issue.Fields)
	json.Unmarshal(fieldsJson, &fields)
	// Process custom fields
	issueTypeName := fields.Issuetype.Name
	issueCustomFields := (*customFieldsByIssueType)[issueTypeName]
	for _, customField := range issueCustomFields {
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
		} else if customField.Custom == sprintCustomType {
			if customField.Name == "Sprint" {
				sprints := []jiramodel.IssueSprint{}
				sprintsJson, _ := json.Marshal(fieldVal)
				json.Unmarshal(sprintsJson, &sprints)
				fields.Sprints = sprints
			}
		}
	}
	issueState, _ := database.SaveIssue(poll.ID, issue, &fields, parentId)

	if poll.HeadIssueStateID == 0 {
		poll.HeadIssueStateID = issueState.ID
		database.UpdatePoll(poll)
	}

	return issueState.IssueID, fields.Subtasks
}
