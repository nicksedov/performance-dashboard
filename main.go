package main

import (
	"flag"
	"fmt"

	"performance-dashboard/pkg/handler"
	"performance-dashboard/pkg/jira"
	"performance-dashboard/pkg/model"
	"performance-dashboard/pkg/profiles"
)

func main() {
	flag.Parse()
	
	config := profiles.GetSettings()
	boardId := config.JiraConfig.BoardID
	projectKey := config.JiraConfig.ProjectKey

	// Get active sprint
	var sprintApiHandler handler.ResponseHandler[model.Sprint] = &handler.SprintHandler{}
	getSprintApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint?state=active", boardId)
	sprint := jira.Query[model.Sprint]("GET", getSprintApiPath, &sprintApiHandler)
	fmt.Printf("Active sprint ID: %d", sprint.ID)

	// Get project info
	// Get active sprint
	var projectApiHandler handler.ResponseHandler[model.Project] = &handler.ProjectHandler{}
	getProjectApiPath := fmt.Sprintf("/rest/api/3/project/%s", projectKey)
	project := jira.Query[model.Project]("GET", getProjectApiPath, &projectApiHandler)
	fmt.Printf("Active project key: %s", project.Key)
	

}
