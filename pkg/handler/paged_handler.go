package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"performance-dashboard/pkg/model"

	"github.com/mitchellh/mapstructure"
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

	data := &model.Pagination{}
	err = json.Unmarshal(body, data)
	if err != nil {
		h.OnError("Error parsing JSON", err)
	}

	mapstructure.Decode(data.Values[0], dto)
}

func (PagedHandler[T]) OnError(reason string, e error) {
	log.Printf("%s:\n%v\n", reason, e)
}
