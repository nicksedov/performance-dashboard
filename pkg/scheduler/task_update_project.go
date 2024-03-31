package scheduler

import (
	"fmt"
	"performance-dashboard/pkg/handler"
	"performance-dashboard/pkg/jira"
	"performance-dashboard/pkg/model"
	"performance-dashboard/pkg/profiles"
)

func updateProject() error {

	config := profiles.GetSettings()
	projectKey := config.JiraConfig.ProjectKey

	// Get project info
	var projectApiHandler handler.ResponseHandler[model.Project] = &handler.ProjectHandler{}
	getProjectApiPath := fmt.Sprintf("/rest/api/3/project/%s", projectKey)
	project := jira.Query[model.Project]("GET", getProjectApiPath, &projectApiHandler)
	fmt.Printf("Active project key: %s\n", project.Key)
	return nil
} 