package scheduler

import (
	"encoding/json"
	"fmt"
	"log"
	"performance-dashboard/pkg/database"
	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/jira/http"
	"performance-dashboard/pkg/jira/model"
	"performance-dashboard/pkg/profiles"

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
	sprints := jira.QueryPaged("GET", getSprintApiPath, &[]model.Sprint{})

	var sprintIds map[int]bool = make(map[int]bool, len(*sprints))
	for _, sprint := range *sprints {
		if sprint.State == "active" {
			log.Printf("Active sprint found: '%s'\n", sprint.Name)
			sprintIds[sprint.ID] = false
		} else if sprint.State == "closed" {
			if database.CompletionPollRequired(sprint.ID) {
				sprintIds[sprint.ID] = true
			}
		}
		database.SaveSprint(&sprint)
	}
	if len(sprintIds) == 0 {
		log.Println("No sprints selected for processing, bypassing issue states poll")
		return nil
	}

	// Polling issue states for
	for sprintID, isCompletionPoll := range sprintIds {
		poll, _ := database.NewPoll(sprintID)
		log.Printf("Collecting issue statuses for sprint with ID '%d'\n", sprintID)
		getIssuesApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint/%d/issue?maxResults=300", boardId, sprintID)
		issues := jira.QueryOne("GET", getIssuesApiPath, &model.Issues{})
		customFieldsByIssueType := getCustomFields()

		for _, issue := range issues.Issues {
			issueId, subtasks := saveIssueState(poll, &issue, customFieldsByIssueType, 0)
			for _, subtask := range subtasks {
				subtaskDetails := jira.QueryOne("GET", subtask.Self, &model.Issue{})
				saveIssueState(poll, subtaskDetails, customFieldsByIssueType, issueId)
			}
		}
		database.CommitPoll(poll)
		database.UpdateSprintPoll(poll.ActiveSprint, poll.ID, isCompletionPoll)
	}
	return nil
}

func getCustomFields() *map[string][]dto.IssueMetadata {
	customFields, _ := database.Read[dto.IssueMetadata](
		func(items *[]dto.IssueMetadata, db *gorm.DB) {
			db.Where("key like 'customfield_%'").Find(items)
		})
	customFieldsByIssueType := make(map[string][]dto.IssueMetadata)
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
				customFieldsByIssueType[key] = make([]dto.IssueMetadata, 0, len(*customFields))
			}
			cf := customFieldsByIssueType[key]
			customFieldsByIssueType[key] = append(cf, customField)
		}
	}
	return &customFieldsByIssueType
}

func saveIssueState(poll *dto.Poll, issue *model.Issue, customFieldsByIssueType *map[string][]dto.IssueMetadata, parentId int) (int, []model.Issue) {
	fields := model.IssueFields{}
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
		if customField.Custom == storyPointCustomType || customField.Name == "Story Points" {
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
				sprints := []model.IssueSprint{}
				sprintsJson, _ := json.Marshal(fieldVal)
				json.Unmarshal(sprintsJson, &sprints)
				fields.Sprints = sprints
			}
		}
	}
	issueState := database.SaveIssue(poll.ID, issue, &fields, parentId)

	if poll.HeadIssueStateID == 0 {
		poll.HeadIssueStateID = issueState.ID
	}

	return issueState.IssueID, fields.Subtasks
}
