package database

import (
	"encoding/json"
	"fmt"
	"log"
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
		dbConfigs := profiles.GetSettings().DbConfig
		dbConfig := dbConfigs.DbNode[0]
		postgresCfg, err := buildPostgresConfig(dbConfig)
		if err != nil {
			return err
		}
		db, err = gorm.Open(postgres.New(*postgresCfg), buildGormConfig(dbConfig))
		if err != nil {
			return err
		} else {
			log.Printf("Database connection established to postgres://%s:%d/%s\n", 
				dbConfig.Host, dbConfig.Port, dbConfig.DbName)
		}
		sqlDb, err := db.DB()
		if err != nil {
			return err
		}
		sqlDb.SetMaxIdleConns(2)
		sqlDb.SetMaxOpenConns(2)
		dbStats, err := json.MarshalIndent(sqlDb.Stats(), "", "  ")
		if err != nil {
			return err
		}
		log.Printf("Database connection information: %s", string(dbStats))
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
			return tx.Error
		}
	}
	return err
}

func buildPostgresConfig(dbConfig profiles.Database) (*postgres.Config, error) {
	dsnFormat := "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s"
	dsn := fmt.Sprintf(dsnFormat,
		dbConfig.Host, dbConfig.Port, dbConfig.DbName, 
		dbConfig.User, dbConfig.Password, dbConfig.SSLMode)
	postgresCfg := &postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	return postgresCfg, nil
}

func buildGormConfig(dbConfig profiles.Database) *gorm.Config {
	gormCfg := &gorm.Config{
		Logger:      logger.Default,
	}
	if dbConfig.SearchPath != "" {
		searchPathNamingStrategy := schema.NamingStrategy{
			TablePrefix: dbConfig.SearchPath + ".",
		}
		gormCfg.NamingStrategy = searchPathNamingStrategy
	}
	return gormCfg
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
