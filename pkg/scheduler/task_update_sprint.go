package scheduler

import (
	//"encoding/json"
	"fmt"
	"log"
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

	log.Printf("Number of issues in active Sprint [%d]: %d\n", sprint.ID, len(issues.Issues))
	for _, issue := range issues.Issues {
		log.Printf("  - [%s] %s %s", issue.Key, issue.Fields.Issuetype.Name, issue.Fields.Summary)
	}

	//issuesJson, _ := json.MarshalIndent(issues, "", "  ")
	//fmt.Printf("%s\n", string(issuesJson))

	return nil
} 