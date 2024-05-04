package database

import (
	"fmt"
	"time"

	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/profiles"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitializeDB() error {
	var err error
	if db == nil {
		dbConfig := profiles.GetSettings().DbConfig
		dsnFormat := "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s"
		dsn := fmt.Sprintf(dsnFormat,
			dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.User, dbConfig.Password, dbConfig.SSLMode)
		gormCfg := &gorm.Config{
			PrepareStmt: false,
			Logger:      logger.Default,
		}
		if dbConfig.SearchPath != "" {
			searchPathNamingStrategy := schema.NamingStrategy{
				TablePrefix: dbConfig.SearchPath + ".",
			}
			gormCfg.NamingStrategy = searchPathNamingStrategy
		}
		postgresCfg := postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}
		db, err = gorm.Open(postgres.New(postgresCfg), gormCfg)
		if err != nil {
			return err
		}
		err = db.AutoMigrate(
			&dto.ApplicationLog{},
			&dto.IssueMetadata{},
			&dto.Sprint{},
			&dto.SprintPoll{},
			&dto.Account{},
			&dto.Poll{},
			&dto.Issue{},
			&dto.IssueState{},
			&dto.IssueSprint{},
			&dto.IssueAssigneeTransitions{},
		)
		if err == nil {
			tx := db.Save(&dto.ApplicationLog{Timestamp: time.Now(), Log: "Database connection created"})
			err = tx.Error
		}
	}
	return err
}

func Read[T any](selector func(items *[]T, db *gorm.DB)) (*[]T, error) {
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
