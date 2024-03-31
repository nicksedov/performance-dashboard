package jira

import (
	"fmt"
	"net/http"
	"performance-dashboard/pkg/profiles"
	"performance-dashboard/pkg/handler"
)

var client *http.Client

func getClient() *http.Client {
	if client == nil {
		client = &http.Client{}
	}
	return client
}

func Query[T any](apiMethod string, apiPath string, respHandler *handler.ResponseHandler[T]) *T {

	// Create HTTP client
	c := getClient()

	// Prepare request
	settings := profiles.GetSettings()
	queryPath := fmt.Sprintf("%s%s", settings.JiraConfig.BaseURL, apiPath)
	req, err := http.NewRequest(apiMethod, queryPath, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v", err)
		return nil
	}

	// Set authentication headers
	req.SetBasicAuth(settings.JiraConfig.Auth.ClientId, settings.JiraConfig.Auth.ApiToken)

	// Make request
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v", err)
		return nil
	}

	return (*respHandler).Handle(resp)
}
