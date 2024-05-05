package database

import (
	"log"
	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/jira/model"
)

func SaveIssueMetadata(f *model.IssueFieldMeta, issueTypeName string, untranslatedName string) {
	var key string
	if f.Key != "" {
		key = f.Key
	} else {
		key = f.FieldID
	}
	newIssueMetadata := dto.IssueMetadata{
		Name:             f.Name,
		Key:              key,
		Type:             f.Schema.Type,
		Custom:           f.Schema.Custom,
		IssueTypeName:    issueTypeName,
		UntranslatedName: untranslatedName,
	}
	existing := dto.IssueMetadata{}
	tx := db.Where(dto.IssueMetadata{Name: f.Name, IssueTypeName: issueTypeName}).First(&existing)
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
}
