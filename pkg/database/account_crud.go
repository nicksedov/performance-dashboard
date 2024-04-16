package database

import (
	"errors"
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
)

func SaveAccount(actor *jira.RoleActor, role string) error {
	_, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return err
	}
	if actor.ActorUser.AccountID != "" {
		account := &database.Account{
			ID:           actor.ActorUser.AccountID,
			AccountType:  actor.Type,
			Role:         role,
			DisplayName:  actor.DisplayName,
		}
		db.Save(account)
	} else {
		return errors.New("invalid accoint ID")
	}

	return nil
}
