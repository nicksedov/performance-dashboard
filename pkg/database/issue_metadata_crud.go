package database

import (
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
)

func SaveIssueMetadata(f *jira.IssueFieldMeta, issueTypeName string, untranslatedName string) error {
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
		IssueTypeName: issueTypeName,
		UntranslatedName: untranslatedName,
	}
	existing := database.IssueMetadata{}
	tx := db.Where(database.IssueMetadata{Name: f.Name, IssueTypeName: issueTypeName}).First(&existing)
	if tx.Error == nil {
		newIssueMetadata.ID = existing.ID
		if existing != newIssueMetadata {
			db.Save(&newIssueMetadata)
		} else {
			log.Printf("Issue field '%s' of type '%s' is already known\n", f.Name, issueTypeName)
		}
	} else {
		db.Save(&newIssueMetadata)
	}
	return nil
}
