package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"performance-dashboard/pkg/jira/model"
)

type PagedHandler[T any] struct {
}

func (h PagedHandler[T]) Handle(resp *http.Response, dto *T) {

	// Prepare request
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		h.OnError("Error reading response body", err)
	}

	data := &jira.Pagination{}
	err = json.Unmarshal(body, data)
	if err != nil {
		h.OnError("Error parsing JSON", err)
	}
	if len(data.Values) > 0 {
		value := data.Values[0]
		jsonListItem, _ := json.Marshal(value)
		json.Unmarshal(jsonListItem, dto)
	}
}

func (PagedHandler[T]) OnError(reason string, e error) {
	log.Printf("%s:\n%v\n", reason, e)
}