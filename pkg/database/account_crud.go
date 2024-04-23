package database

import (
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

	account := &database.Account{
		ID:           actor.ID,
		AccountID:    actor.ActorUser.AccountID,
		AccountType:  actor.Type,
		Role:         role,
		DisplayName:  actor.DisplayName,
		EmailAddress: actor.EmailAddress,
	}
	db.Save(account)

	return nil
}
