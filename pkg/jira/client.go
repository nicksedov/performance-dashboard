package jira

import (
	"log"
	"net/http"
	"net/url"
	"performance-dashboard/pkg/handler"
	"performance-dashboard/pkg/profiles"
	"strings"
)

var client *http.Client

func Query[T any](apiMethod string, apiPath string, respHandler *handler.ResponseHandler[T]) *T {

	// Create HTTP client
	c := getClient()

	// Prepare request
	settings := profiles.GetSettings()
	queryPath := buildUrl(settings.JiraConfig.BaseURL, apiPath)
	req, err := http.NewRequest(apiMethod, queryPath, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil
	}

	// Set authentication headers
	req.SetBasicAuth(settings.JiraConfig.Auth.ClientId, settings.JiraConfig.Auth.ApiToken)

	// Make request
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil
	}

	return (*respHandler).Handle(resp)
}

func getClient() *http.Client {
	if client == nil {
		client = &http.Client{}
	}
	return client
}

func buildUrl(baseUrl, apiPath string) string {
	if (strings.HasPrefix(apiPath, "http:")  || strings.HasPrefix(apiPath, "https:")) {
		return apiPath
	} else {
		result, _ := url.JoinPath(baseUrl, apiPath)
		return result
	}
}
