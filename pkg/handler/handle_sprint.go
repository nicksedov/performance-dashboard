package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"performance-dashboard/pkg/jira"
	"performance-dashboard/pkg/model"

	"github.com/mitchellh/mapstructure"
)

type SprintHandler struct {
	result *model.Sprint
}

func (h SprintHandler) Handle(resp *http.Response) *model.Sprint {

	// Prepare request
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		h.OnError("Error reading response body", err)
		return nil
	}

	data := &model.Pagination{}
	err = json.Unmarshal(body, data)
	if err != nil {
		h.OnError("Error parsing JSON", err)
		return nil
	}

	h.result = &model.Sprint{}
	mapstructure.Decode(data.Values[0], h.result)

	return h.result
}

func (SprintHandler) OnError(reason string, e error) {
	fmt.Printf("%s:\n%v\n", reason, e)
}

func (SprintHandler) Inst() jira.ResponseHandler[model.Sprint] {
	return &SprintHandler{}
}
