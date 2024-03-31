package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"performance-dashboard/pkg/model"
)

type RoleHandler struct {
	result *model.Role
}

func (h RoleHandler) Handle(resp *http.Response) *model.Role {
	// Prepare request
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		h.OnError("Error reading response body", err)
		return nil
	}

	h.result = &model.Role{}
	err = json.Unmarshal(body, h.result)
	if err != nil {
		h.OnError("Error parsing JSON", err)
		return nil
	}

	return h.result
}

func (RoleHandler) OnError(reason string, e error) {
	log.Printf("%s:\n%v\n", reason, e)
}
