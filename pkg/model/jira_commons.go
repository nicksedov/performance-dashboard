package model

type Pagination struct {
	MaxResults int              `json:"maxResults"`
	StartAt    int              `json:"startAt"`
	Total      int              `json:"total"`
	IsLast     bool             `json:"isLast"`
	Values     []map[string]any `json:"values"`
}

type AvatarUrls struct {
	Size16X16   string `json:"16x16"`
	Size24X24   string `json:"24x24"`
	Size32X32   string `json:"32x32"`
	Size48X48   string `json:"48x48"`
}