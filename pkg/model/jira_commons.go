package model

type Pagination struct {
	MaxResults int              `json:"maxResults"`
	StartAt    int              `json:"startAt"`
	Total      int              `json:"total"`
	IsLast     bool             `json:"isLast"`
	Values     []map[string]any `json:"values"`
}
