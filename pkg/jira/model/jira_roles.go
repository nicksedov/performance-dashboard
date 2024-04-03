package jira

type Role struct {
	Actors []struct {
		ActorGroup struct {
			DisplayName string `json:"displayName"`
			GroupID     string `json:"groupId"`
			Name        string `json:"name"`
		} `json:"actorGroup,omitempty"`
		DisplayName string `json:"displayName"`
		ID          int    `json:"id"`
		Name        string `json:"name,omitempty"`
		Type        string `json:"type"`
		User        string `json:"user,omitempty"`
		ActorUser   struct {
			AccountID string `json:"accountId"`
		} `json:"actorUser,omitempty"`
	} `json:"actors"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Scope       struct {
		Project struct {
			ID   string `json:"id"`
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"project"`
		Type string `json:"type"`
	} `json:"scope"`
	Self string `json:"self"`
}
