package scheduler

import (
	"encoding/json"
	"fmt"
	"performance-dashboard/pkg/jira"
	"performance-dashboard/pkg/model"
	"performance-dashboard/pkg/profiles"
	"performance-dashboard/pkg/util"
)

func updateProject() error {

	config := profiles.GetSettings()
	projectKey := config.JiraConfig.ProjectKey

	// Get project info
	getProjectApiPath := fmt.Sprintf("/rest/api/2/project/%s", projectKey)
	project := jira.QueryOne("GET", getProjectApiPath, &model.Project{})
	projectJson, _ := json.Marshal(project)
	fmt.Printf("Active project key: %s\nProject Details:\n", project.Key)
	util.PrettyPrintJSON(projectJson)

	roles := project.Roles

	for _, getRoleApi := range roles {
		role := jira.QueryOne("GET", getRoleApi, &model.Role{})
		fmt.Printf("Project %s has role %s with %d actors\n", projectKey, role.Name, len(role.Actors))
		for _, actor := range role.Actors {
			fmt.Printf("  - %s [%d]\n", actor.DisplayName, actor.ID)
		}
	}

	return nil
}
