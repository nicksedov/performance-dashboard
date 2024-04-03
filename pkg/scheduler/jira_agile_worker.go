package scheduler

import (
	"fmt"
	"log"
	"performance-dashboard/pkg/database"
	"performance-dashboard/pkg/jira/http"
	model "performance-dashboard/pkg/jira/model"
	"performance-dashboard/pkg/profiles"
)

func jiraAgileWorker() error {

	config := profiles.GetSettings()
	boardId := config.JiraConfig.BoardID

	// Get active sprint
	getSprintApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint?state=active", boardId)
	sprint := jira.QueryPaged("GET", getSprintApiPath, &model.Sprint{})

	log.Printf("Active Sprint: %d\n", sprint.ID)

	database.SaveSprint(sprint)

	getIssuesApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint/%d/issue", boardId, sprint.ID)
	issues := jira.QueryOne("GET", getIssuesApiPath, &model.Issues{})

	log.Println("Sprint issues:")
	for _, issue := range issues.Issues {
		log.Printf("  - [%s] %s %s", issue.Key, issue.Fields.Issuetype.Name, issue.Fields.Summary)
	}

	return nil
}
