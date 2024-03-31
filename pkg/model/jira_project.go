package model

type Project struct {
	ProjectKey string   `json:"key"`
	Members    []Member `json:"members"`
}

type Member struct {
	Name     string `json:"displayName"`
	Email    string `json:"emailAddress"`
	Username string `json:"accountId"`
}

