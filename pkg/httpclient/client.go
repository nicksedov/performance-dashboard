package httpclient

import (
	"log"
	"net/http"
	"net/url"
	"performance-dashboard/pkg/jira/handler"
	"performance-dashboard/pkg/profiles"
	"strconv"
	"time"
)

var (
	client *http.Client
	lastRequestTimestamp time.Time
	requestRateLimit int = 8
	retryLimit int = 3
	interval time.Duration
)


func QueryOne[T any](apiMethod string, apiPath string, dto *T) *T {
	resp := httpQuery(apiMethod, apiPath)
	if resp != nil {
		respHandler := handler.SimpleHandler[T]{}
		respHandler.Handle(resp, dto)
	}
	return dto
}

func QueryPaged[T any](apiMethod string, apiPath string, dto *[]T) *[]T {
	resp := httpQuery(apiMethod, apiPath)
	if resp != nil {
		defer resp.Body.Close()
		respHandler := handler.PagedHandler[T]{}
		respHandler.Handle(resp, dto)
	}
	return dto
}

func getClient() *http.Client {
	if client == nil {
		client = &http.Client{}
		settings := profiles.GetSettings()
		client.Timeout = settings.HttpClientConfig.RequestTimeout
		if settings.HttpClientConfig.RequestRateLimit > 0 {
			requestRateLimit = settings.HttpClientConfig.RequestRateLimit
		}
		if settings.HttpClientConfig.RetryLimit > 0 {
			retryLimit = settings.HttpClientConfig.RetryLimit
		}
		interval = time.Second / time.Duration(requestRateLimit) 
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

func httpQuery(apiMethod string, apiPath string) *http.Response {

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
	resp, err := doRequest(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil
	}
	return resp
}

func doRequest(req *http.Request) (*http.Response, error) {
	return doRetryableRequest(req, 0)
}

func doRetryableRequest(req *http.Request, retryCount int) (*http.Response, error) {

	c:= getClient()

	// Adding pause to not exceed given RequestRateLimit or extended pause on retry attempt 
	timePassed := time.Since(lastRequestTimestamp)
	awaitDuration := interval * time.Duration(retryCount + 1) 
	if awaitDuration > timePassed {
		time.Sleep(awaitDuration - timePassed)
		lastRequestTimestamp = time.Now()
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
    // Retry on "Too many requests" or "Service Unavailable"
	if (resp.StatusCode == 429 || resp.StatusCode == 503) && retryCount < retryLimit {
		log.Printf("Warning: HTTP request '%s %s' returned status '%s', retrying...", req.Method, req.RequestURI, resp.Status)
		if len(resp.Header["Retry-After"]) > 0 {
			onRetryAfter(resp.Header["Retry-After"][0])
		}
		return doRetryableRequest(req, retryCount + 1)
	}
	
	return resp, nil
}

func onRetryAfter(retryAfter string) {
	var sleepDuration time.Duration
	sec, err := strconv. Atoi(retryAfter) 
	if err == nil {
		sleepDuration = time.Duration(sec) * time.Second
	} else {
		timestamp, err := time.Parse(time.RFC1123, retryAfter)
		if err == nil {
			sleepDuration = time.Until(timestamp)
		}
	}
	if sleepDuration > 0 {
		log.Printf("Setting retry interval to '%v' according to 'Retry-After' header value", sleepDuration)
		time.Sleep(sleepDuration)
	}
}