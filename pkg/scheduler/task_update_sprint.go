package scheduler

import (
	"fmt"
	"performance-dashboard/pkg/handler"
	"performance-dashboard/pkg/jira"
	"performance-dashboard/pkg/model"
	"performance-dashboard/pkg/profiles"
)

func updateSprint() error {

	config := profiles.GetSettings()
	boardId := config.JiraConfig.BoardID

	// Get active sprint
	var sprintApiHandler handler.ResponseHandler[model.Sprint] = &handler.SprintHandler{}
	getSprintApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint?state=active", boardId)
	sprint := jira.Query[model.Sprint]("GET", getSprintApiPath, &sprintApiHandler)
	fmt.Printf("Active sprint ID: %d\n", sprint.ID)
	return nil
} 