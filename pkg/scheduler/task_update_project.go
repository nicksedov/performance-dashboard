package scheduler

import (
	"encoding/json"
	"fmt"
	"log"
	"performance-dashboard/pkg/jira"
	"performance-dashboard/pkg/model"
	"performance-dashboard/pkg/profiles"
	"performance-dashboard/pkg/util"
	"time"
)

func updateProject() error {

	config := profiles.GetSettings()
	projectKey := config.JiraConfig.ProjectKey

	// Get project info
	getProjectApiPath := fmt.Sprintf("/rest/api/2/project/%s", projectKey)
	project := jira.QueryOne("GET", getProjectApiPath, &model.Project{})
	projectJson, _ := json.Marshal(project)
	
	log.Printf("[%s] Project update\nActive project key: %s\nProject Details:\n", time.Now().Format("01-02-2006 15:04:05"), project.Key)
	util.PrettyPrintJSON(projectJson)

	roles := project.Roles
	for _, getRoleApi := range roles {
		role := jira.QueryOne("GET", getRoleApi, &model.Role{})
		log.Printf("Role lookup URL: %s\nProject %s has role %s with %d actors:\n", getRoleApi, projectKey, role.Name, len(role.Actors))
		for _, actor := range role.Actors {
			log.Printf("  - %s [%d]\n", actor.DisplayName, actor.ID)
		}
	}

	return nil
}
