package model

type Project struct {
	AssigneeType string   `json:"assigneeType"`
	AvatarUrls AvatarUrls `json:"avatarUrls"`
	Components []struct {
		Ari      string `json:"ari"`
		Assignee struct {
			AccountID   string `json:"accountId"`
			AccountType string `json:"accountType"`
			Active      bool   `json:"active"`
			AvatarUrls  AvatarUrls `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Key         string `json:"key"`
			Name        string `json:"name"`
			Self        string `json:"self"`
		} `json:"assignee"`
		AssigneeType        string `json:"assigneeType"`
		Description         string `json:"description"`
		ID                  string `json:"id"`
		IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid"`
		Lead struct {
			AccountID   string `json:"accountId"`
			AccountType string `json:"accountType"`
			Active      bool   `json:"active"`
			AvatarUrls  AvatarUrls `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Key         string `json:"key"`
			Name        string `json:"name"`
			Self        string `json:"self"`
		} `json:"lead"`
		Metadata struct {
			Icon string `json:"icon"`
		} `json:"metadata"`
		Name         string `json:"name"`
		Project      string `json:"project"`
		ProjectID    int    `json:"projectId"`
		RealAssignee struct {
			AccountID   string `json:"accountId"`
			AccountType string `json:"accountType"`
			Active      bool   `json:"active"`
			AvatarUrls  AvatarUrls `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Key         string `json:"key"`
			Name        string `json:"name"`
			Self        string `json:"self"`
		} `json:"realAssignee"`
		RealAssigneeType string `json:"realAssigneeType"`
		Self             string `json:"self"`
	} `json:"components"`
	Description string `json:"description"`
	Email       string `json:"email"`
	ID          string `json:"id"`
	Insight     struct {
		LastIssueUpdateTime string `json:"lastIssueUpdateTime"`
		TotalIssueCount     int    `json:"totalIssueCount"`
	} `json:"insight"`
	IssueTypes []struct {
		AvatarID       int    `json:"avatarId"`
		Description    string `json:"description"`
		HierarchyLevel int    `json:"hierarchyLevel"`
		IconURL        string `json:"iconUrl"`
		ID             string `json:"id"`
		Name           string `json:"name"`
		Self           string `json:"self"`
		Subtask        bool   `json:"subtask"`
		EntityID       string `json:"entityId,omitempty"`
		Scope          struct {
			Project struct {
				ID string `json:"id"`
			} `json:"project"`
			Type string `json:"type"`
		} `json:"scope,omitempty"`
	} `json:"issueTypes"`
	Key  string `json:"key"`
	Lead struct {
		AccountID   string `json:"accountId"`
		AccountType string `json:"accountType"`
		Active      bool   `json:"active"`
		AvatarUrls AvatarUrls `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Key         string `json:"key"`
		Name        string `json:"name"`
		Self        string `json:"self"`
	} `json:"lead"`
	Name            string `json:"name"`
	ProjectCategory struct {
		Description string `json:"description"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Self        string `json:"self"`
	} `json:"projectCategory"`
	Properties map[string]string `json:"properties"`
	Roles      map[string]string `json:"roles"`
	Self       string `json:"self"`
	Simplified bool   `json:"simplified"`
	Style      string `json:"style"`
	URL        string `json:"url"`
	Versions   []string  `json:"versions"`
}