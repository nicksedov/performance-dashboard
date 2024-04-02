package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SimpleHandler[T any] struct {
}

func (h SimpleHandler[T]) Handle(resp *http.Response, dto *T) {
	// Prepare request
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		h.OnError("Error reading response body", err)
	}
	
	err = json.Unmarshal(body, dto)
	if err != nil {
		errorMsg := fmt.Sprintf("Expected valid JSON in HTTP response\n  HTTP request: %s %s\n HTTP responded: %s", 
			resp.Request.Method, resp.Request.RequestURI, string(body))
		h.OnError(errorMsg, err)
	}
}

func (SimpleHandler[T]) OnError(reason string, e error) {
	log.Printf("%s:\n%v\n", reason, e)
}
