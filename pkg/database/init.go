package database

import (
	"fmt"
	"log"

	database "performance-dashboard/pkg/database/model"
	"performance-dashboard/pkg/profiles"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb() (*gorm.DB, error) {
	var err error
	if db == nil {
		dbConfig := profiles.GetSettings().DbConfig
		var dsn string
		dsnFormat := "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s"
		dsn = fmt.Sprintf(dsnFormat,
			dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.User, dbConfig.Password, dbConfig.SSLMode)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbConfig.SearchPath != "" {
			log.Printf("Switching database schema to '%s'\n", dbConfig.SearchPath)
			db.Exec("set search_path to ?", dbConfig.SearchPath)
		}
		db.AutoMigrate(
			&database.IssueMetadata{},
			&database.Sprint{},
			&database.Account{},
			&database.Poll{},
			&database.Issue{},
			&database.IssueState{},
			&database.IssueClosedSprint{},
		)
	}
	return db, err
}

func Read[T any](selector func(items *[]T, db *gorm.DB)) (*[]T, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	items := new([]T)
	selector(items, db)
	return items, nil
}

func GetAll[T any]() (*[]T, error) {
	selectAll := func(items *[]T, db *gorm.DB) {
		db.Order("id").Find(items)
	}
	return Read(selectAll)
}
