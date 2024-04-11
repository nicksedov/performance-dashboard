package database

import (
	"errors"
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
)

func SaveAccount(acc *jira.Account) error {
	_, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return err
	}
	if acc.AccountID != "" {
		account := &database.Account{
			ID:           acc.AccountID,
			AccountType:  acc.AccountType,
			Active:       acc.Active,
			DisplayName:  acc.DisplayName,
			EmailAddress: acc.EmailAddress,
			Key:          acc.Key,
			Name:         acc.Name,
			TimeZone:     acc.TimeZone,
		}
		db.Save(account)
	} else {
		return errors.New("invalid accoint ID")
	}

	return nil
}
