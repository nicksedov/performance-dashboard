package scheduler

import (
	"fmt"
	"performance-dashboard/pkg/jira"
	"performance-dashboard/pkg/model"
	"performance-dashboard/pkg/profiles"
)

func updateSprint() error {

	config := profiles.GetSettings()
	boardId := config.JiraConfig.BoardID

	// Get active sprint
	getSprintApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint?state=active", boardId)
	sprint := jira.QueryPaged("GET", getSprintApiPath, &model.Sprint{})
	
	getIssuesApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint/%d/issue", boardId, sprint.ID)
	issues := jira.QueryOne("GET", getIssuesApiPath, &model.Issues{})

	fmt.Printf("Active sprintIssues in Active Sprint []:")
	for _, issue := range issues.Issues {
		fmt.Printf("\nKey: %s\n", issue.GetKey())
		for k, v := range *issue.GetAllParams() {
			fmt.Printf("  %s: %+v\n", k, v)
		}
	}

	return nil
} 