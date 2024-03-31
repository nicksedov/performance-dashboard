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
	var sprintApiHandler jira.ResponseHandler[model.Sprint] = &handler.SprintHandler{}

	getSprintApiPath := fmt.Sprintf("/rest/agile/1.0/board/%s/sprint?state=active", boardId)
	sprint := jira.Query[model.Sprint]("GET", getSprintApiPath, &sprintApiHandler)

	sprintId := sprint.ID

	fmt.Printf("Active sprint ID: %d", sprintId)
}
