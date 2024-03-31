package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"performance-dashboard/pkg/model"
)

type ProjectHandler struct {
	result *model.Project
}

func (h ProjectHandler) Handle(resp *http.Response) *model.Project {
	// Prepare request
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		h.OnError("Error reading response body", err)
		return nil
	}

	h.result = &model.Project{}
	err = json.Unmarshal(body, h.result)
	if err != nil {
		h.OnError("Error parsing JSON", err)
		return nil
	}

	return h.result
}

func (ProjectHandler) OnError(reason string, e error) {
	fmt.Printf("%s:\n%v\n", reason, e)
}
