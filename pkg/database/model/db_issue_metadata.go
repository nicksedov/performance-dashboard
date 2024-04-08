package database

type IssueMetadata struct {
	ID            int        `gorm:"unique;primaryKey"`
	Key           string
	Name          string
	Type          string
	Custom        string
}

