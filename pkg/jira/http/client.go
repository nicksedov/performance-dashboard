package jira

import (
	"log"
	"net/http"
	"net/url"
	"performance-dashboard/pkg/jira/handler"
	"performance-dashboard/pkg/profiles"
)

var client *http.Client

func QueryOne[T any](apiMethod string, apiPath string, dto *T) *T {
	resp := queryRaw(apiMethod, apiPath)
	if resp != nil {
		respHandler := handler.SimpleHandler[T]{}
		respHandler.Handle(resp, dto)
	}
	return dto
}

func QueryPaged[T any](apiMethod string, apiPath string, dto *T) *T {
	resp := queryRaw(apiMethod, apiPath)
	if resp != nil {
		respHandler := handler.PagedHandler[T]{}
		respHandler.Handle(resp, dto)
	}
	return dto
}

func getClient() *http.Client {
	if client == nil {
		client = &http.Client{}
	}
	return client
}

func buildUrl(baseUrl, apiPath string) string {
	u, err := url.Parse(apiPath)
    if err != nil {
        log.Fatal(err)
    }
	if u.Scheme != "" && u.Host != "" {
		return u.String()
	} else {
		base, err := url.Parse(baseUrl)
		if err != nil {
			log.Fatal(err)
		}
		return base.ResolveReference(u).String()
	}
}

func queryRaw(apiMethod string, apiPath string) *http.Response {

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
	authSettings := settings.JiraConfig.Auth
	if authSettings.Type == "basic" {
		req.SetBasicAuth(settings.JiraConfig.Auth.ClientId, settings.JiraConfig.Auth.ApiToken)
	} else if authSettings.Type == "apitoken" {
		req.Header.Set("Authorization", "Bearer " + settings.JiraConfig.Auth.ApiToken)
	} else {
		log.Printf("Warning: invalid authorization type %s", authSettings.Type)
	}

	// Make request
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil
	}

	return resp
}
