package database

import (
	"log"
	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/jira/model"
)

func SaveAccount(actor *model.RoleActor, role string) {
	newAccount := dto.Account{
		ID:           actor.ID,
		AccountID:    actor.ActorUser.AccountID,
		AccountType:  actor.Type,
		Role:         role,
		DisplayName:  actor.DisplayName,
		EmailAddress: actor.EmailAddress,
	}
	existing := dto.Account{}
	tx := db.Where(dto.Account{Role: role, DisplayName: newAccount.DisplayName}).First(&existing)
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
