package scheduler

import (
	"fmt"
	"log"
	"performance-dashboard/pkg/database"
	jira "performance-dashboard/pkg/jira/http"
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
		time.Sleep(200 * time.Millisecond)
		role := jira.QueryOne("GET", getRoleApi, &model.Role{})
		log.Printf("Project %s has role %s with %d actors\n", projectKey, role.Name, len(role.Actors))
		for _, actor := range role.Actors {
			database.SaveAccount(&actor, role.Name)
		}
	}

	getIssueFieldsApiPath := fmt.Sprintf("/rest/api/2/issue/createmeta?projectKeys=%s&expand=projects.issuetypes.fields", projectKey)
	issueFields := jira.QueryOne("GET", getIssueFieldsApiPath, &model.IssueFieldsMeta{})
	if len(issueFields.Projects) == 0 || len(issueFields.Projects[0].Issuetypes) == 0 {
		log.Printf("Unable to get metadata for issues of related to project '%s'", projectKey)
		return nil
	}
	projectMeta := issueFields.Projects[0]
	for _,storyMetadata := range projectMeta.Issuetypes {
		storyFields := storyMetadata.Fields
		issueTypeName := storyMetadata.Name
		untranslatedName := storyMetadata.UntranslatedName
		for _, val := range storyFields {
			database.SaveIssueMetadata(&val, issueTypeName, untranslatedName)
		}
	}  

	return nil
}
