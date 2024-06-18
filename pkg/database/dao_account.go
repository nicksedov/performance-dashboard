package database

import (
	"errors"
	"log"
	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/jira/model"

	"gorm.io/gorm"
)

const extAccountLabel string = "[external]"
const extAccountRole string = "External participant"

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
	tx := GetDB().Where(dto.Account{Role: role, DisplayName: newAccount.DisplayName}).First(&existing)
	if tx.Error == nil {
		newAccount.ID = existing.ID
		if existing != newAccount {
			GetDB().Save(&newAccount)
		} else {
			log.Printf("Account '%s' with role '%s' is already known\n", newAccount.DisplayName, role)
		}
	} else {
		GetDB().Save(&newAccount)
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
	tx := GetDB().Where(dto.Account{Role: extAccountRole, DisplayName: newAccount.DisplayName}).First(&existing)
	if tx.Error == nil {
		newAccount.ID = existing.ID
		if existing != newAccount {
			GetDB().Save(&newAccount)
		} else {
			log.Printf("Account '%s' with role '%s' is already known\n", newAccount.DisplayName, extAccountRole)
		}
	} else {
		lastExtAccount := &dto.Account{}
		tx := GetDB().Where("account_type LIKE ?", extAccountLabel+"%").Order("id DESC").First(lastExtAccount)
		if tx.Error == nil {
			newAccount.ID = lastExtAccount.ID + 1
		} else if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			newAccount.ID = -10000
		} else {
			log.Printf("Error saving account record '%s' with role '%s': %v\n", newAccount.DisplayName, extAccountRole, tx.Error)
		}
		if newAccount.ID != 0 {
			log.Printf("Saving account record '%s' with role '%s' and ID=%d\n", newAccount.DisplayName, extAccountRole, newAccount.ID)
			GetDB().Save(&newAccount)
		}

	}
	return &newAccount
}
