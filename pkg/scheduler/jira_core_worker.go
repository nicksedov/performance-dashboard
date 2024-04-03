package scheduler

import (
	"fmt"
	"log"
	"performance-dashboard/pkg/jira/http"
	model "performance-dashboard/pkg/jira/model"
	"performance-dashboard/pkg/profiles"
	"time"
)

func jiraCoreWorker() error {

	config := profiles.GetSettings()
	projectKey := config.JiraConfig.ProjectKey

	// Get project info
	getProjectApiPath := fmt.Sprintf("/rest/api/2/project/%s", projectKey)
	project := jira.QueryOne("GET", getProjectApiPath, &model.Project{})

	roles := project.Roles

	for _, getRoleApi := range roles {
		time.Sleep(500 * time.Millisecond)
		role := jira.QueryOne("GET", getRoleApi, &model.Role{})
		log.Printf("Role lookup URL: %s\nProject %s has role %s with %d actors:\n", getRoleApi, projectKey, role.Name, len(role.Actors))
		for _, actor := range role.Actors {
			log.Printf("  - %s [%d]\n", actor.DisplayName, actor.ID)
		}
	}

	return nil
}