package dto

type IssueMetadata struct {
	ID               int        `gorm:"unique;primaryKey"`
	Key              string
	Name             string
	Type             string
	Custom           string
	IssueTypeName    string
	UntranslatedName string
}

