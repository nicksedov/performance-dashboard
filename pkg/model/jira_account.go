package model

type Account struct {
	AccountID   string `json:"accountId"`
	AccountType string `json:"accountType"`
	Active      bool   `json:"active"`
	AvatarUrls AvatarUrls `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Self        string `json:"self"`
	TimeZone    string `json:"timeZone"`
}
