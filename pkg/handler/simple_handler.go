package handler

import (
	"encoding/json"
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
		h.OnError("Error parsing JSON", err)
	}
}

func (SimpleHandler[T]) OnError(reason string, e error) {
	log.Printf("%s:\n%v\n", reason, e)
}
