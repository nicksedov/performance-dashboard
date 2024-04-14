package database

import (
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
)

func SaveIssueMetadata(f *jira.IssueFieldMeta) error {
	_, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return err
	}
	newIssueMetadata := database.IssueMetadata{
		Name:   f.Name,
		Key:    f.Key,
		Type:   f.Schema.Type,
		Custom: f.Schema.Custom,
	}
	db.Where(database.IssueMetadata{Name: f.Name}).Assign(newIssueMetadata).FirstOrCreate(&newIssueMetadata)
	return nil
}
