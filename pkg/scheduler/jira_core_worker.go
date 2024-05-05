package scheduler

import (
	"fmt"
	"log"
	"performance-dashboard/pkg/database"
	"performance-dashboard/pkg/jira/http"
	"performance-dashboard/pkg/jira/model"
	"performance-dashboard/pkg/profiles"
)

func jiraCoreWorker() error {

	config := profiles.GetSettings()
	projectKey := config.JiraConfig.ProjectKey
	project := getProject(projectKey)

	// Get project roles
	roles := project.Roles
	log.Printf("Collecting information about actors in project '%s'\n", projectKey)
	for _, getRoleApi := range roles {
		role := jira.QueryOne("GET", getRoleApi, &model.Role{})
		actorsCount := len(role.Actors)
		if actorsCount > 0 {
			log.Printf("Found %d actors with role %s\n", actorsCount, role.Name)
			for _, actor := range role.Actors {
				database.SaveAccount(&actor, role.Name)
			}
		}
	}

	// Get metadata of issues related to the project
	getIssueFieldsApiPath := fmt.Sprintf("/rest/api/2/issue/createmeta?projectKeys=%s&expand=projects.issuetypes.fields", projectKey)
	issueFields := jira.QueryOne("GET", getIssueFieldsApiPath, &model.IssueFieldsMeta{})
	if len(issueFields.Projects) == 0 || len(issueFields.Projects[0].Issuetypes) == 0 {
		log.Printf("Unable to get metadata for issues related to project '%s'", projectKey)
		return nil
	}
	projectMeta := issueFields.Projects[0]
	for _, storyMetadata := range projectMeta.Issuetypes {
		storyFields := storyMetadata.Fields
		issueTypeName := storyMetadata.Name
		untranslatedName := storyMetadata.UntranslatedName
		log.Printf("Collecting metadata for issues of type '%s'\n", issueTypeName)
		for _, val := range storyFields {
			database.SaveIssueMetadata(&val, issueTypeName, untranslatedName)
		}
	}

	return nil
}

// Get base project information by key
func getProject(projectKey string) *model.Project {
	log.Printf("Collecting information about project '%s'\n", projectKey)
	getProjectApiPath := fmt.Sprintf("/rest/api/2/project/%s", projectKey)
	return jira.QueryOne("GET", getProjectApiPath, &model.Project{})
}
