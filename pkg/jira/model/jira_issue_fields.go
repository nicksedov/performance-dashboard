package jira

type IssueFieldMeta struct {
	AutoCompleteURL string   `json:"autoCompleteUrl"`
	HasDefaultValue bool     `json:"hasDefaultValue"`
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	FieldID         string   `json:"fieldId"`
	Operations      []string `json:"operations"`
	Required        bool     `json:"required"`
	Schema          struct {
		Items  string `json:"items"`
		System string `json:"system"`
		Type   string `json:"type"`
		Custom string `json:"custom"`
	}
	AllowedValues []map[string]any `json:"allowedValues"`
}

type IssueTypeMetadata struct {
	Description string                    `json:"description"`
	Expand      string                    `json:"expand"`
	Fields      map[string]IssueFieldMeta `json:"fields"`
	IconURL     string                    `json:"iconUrl"`
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Scope       struct {
		Project struct {
			ID string `json:"id"`
		} `json:"project"`
		Type string `json:"type"`
	} `json:"scope"`
	Self             string `json:"self"`
	Subtask          bool   `json:"subtask"`
	UntranslatedName string `json:"untranslatedName"`
}

type IssueFieldsMeta struct {
	Expand   string `json:"expand"`
	Projects []struct {
		AvatarUrls AvatarUrls `json:"avatarUrls"`
		Expand     string     `json:"expand"`
		ID         string     `json:"id"`
		Issuetypes []IssueTypeMetadata `json:"issuetypes"`
		Key  string `json:"key"`
		Name string `json:"name"`
		Self string `json:"self"`
	} `json:"projects"`
}
