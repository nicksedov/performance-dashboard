package database

import (
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
)

func SaveAccount(actor *jira.RoleActor, role string) {
	newAccount := database.Account{
		ID:           actor.ID,
		AccountID:    actor.ActorUser.AccountID,
		AccountType:  actor.Type,
		Role:         role,
		DisplayName:  actor.DisplayName,
		EmailAddress: actor.EmailAddress,
	}
	existing := database.Account{}
	tx := db.Where(database.Account{Role: role, DisplayName: newAccount.DisplayName}).First(&existing)
	if tx.Error == nil {
		newAccount.ID = existing.ID
		if existing != newAccount {
			db.Save(&newAccount)
		} else {
			log.Printf("Account '%s' with role '%s' is already known\n", newAccount.DisplayName, role)
		}
	} else {
		db.Save(&newAccount)
	}
}
