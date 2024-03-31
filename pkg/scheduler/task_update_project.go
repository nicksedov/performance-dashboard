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
	getProjectApiPath := fmt.Sprintf("/rest/api/2/project/%s", projectKey)
	project := jira.Query[model.Project]("GET", getProjectApiPath, &projectApiHandler)
	fmt.Printf("Active project key: %s\n", project.Key)

	roles := project.Roles
	var roleApiHandler handler.ResponseHandler[model.Role] = &handler.RoleHandler{}
	
	for _, getRoleApi := range roles {
		role := jira.Query[model.Role]("GET", getRoleApi, &roleApiHandler)
		fmt.Printf("Project %s has role %s with %d actors\n", projectKey, role.Name, len(role.Actors))
		for _, actor := range role.Actors {
			fmt.Printf("  - %s [%d]\n", actor.DisplayName, actor.ID)
		}
	}
	
	return nil
}
