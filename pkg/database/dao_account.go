package database

import (
	"log"
	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/jira/model"
)

type sequence struct {
	NextId int
}
const extAccountLabel string  = "[external]"
const extAccountRole  string = "External participant"

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

func SaveExternalParticipantAccount(actor *model.Account) *dto.Account {
	
	newAccount := dto.Account{
		AccountID:    actor.AccountID,
		AccountType:  extAccountLabel + actor.AccountType,
		Role:         extAccountRole,
		DisplayName:  actor.DisplayName,
		EmailAddress: actor.EmailAddress,
	}
	existing := dto.Account{}
	tx := db.Where(dto.Account{Role: extAccountRole, DisplayName: newAccount.DisplayName}).First(&existing)
	if tx.Error == nil {
		newAccount.ID = existing.ID
		if existing != newAccount {
			db.Save(&newAccount)
		} else {
			log.Printf("Account '%s' with role '%s' is already known\n", newAccount.DisplayName, extAccountRole)
		}
	} else {
		seq := &sequence{}
		row := db.Select("COALESCE(MAX(id), -10000) + 1 as next_id").Where("account_type LIKE ?", extAccountLabel + "%").Table("accounts").Row()
		err := row.Scan(seq)
		if err == nil {
			newAccount.ID = seq.NextId
			db.Save(&newAccount)
		} else {
			log.Printf("Error saving account record '%s' with role '%s'", newAccount.DisplayName, extAccountRole)
		}
	}
	return &newAccount
}