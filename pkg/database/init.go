package database

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/profiles"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db []*gorm.DB
var activeNode int = -1

func InitializeDB() error {
	if len(db) == 0 {
		dbConfigs := profiles.GetSettings().DbConfig
		for _, dbConfig := range dbConfigs.DbNode {
			postgresCfg, err := buildPostgresConfig(dbConfig)
			if err != nil {
				return err
			}
			dbNode, err := gorm.Open(postgres.New(*postgresCfg), buildGormConfig(dbConfig))
			if err != nil {
				return err
			} else {
				log.Printf("Database connection established to postgres://%s:%d/%s\n",
					dbConfig.Host, dbConfig.Port, dbConfig.DbName)
			}
			sqlDb, err := dbNode.DB()
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
			row := sqlDb.QueryRow("SELECT pg_is_in_recovery()")
			var recovery bool
			if err := row.Scan(&recovery); err != nil {
				return err
			} else {
				db = append(db, dbNode)
				if recovery {
					log.Printf("Database is in standby/recovery mode")
				} else {
					activeNode = len(db) - 1
					log.Printf("Database is active")
				} 
			}
		}
	}
	if GetDB() == nil {
		return errors.New("Database not available")
	}
	return autoMigrate()
}

func autoMigrate() error {
	err := GetDB().AutoMigrate(
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
		tx := GetDB().Save(&dto.ApplicationLog{Timestamp: time.Now(), Log: "Database connection created"})
		return tx.Error
	}
	return err
}

func GetDB() *gorm.DB {
	if activeNode > -1 && activeNode < len(db) {
		return db[activeNode]
	}
	return nil
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
		Logger: logger.Default,
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
	selector(items, GetDB())
	return items, nil
}

func GetAll[T any]() (*[]T, error) {
	selectAll := func(items *[]T, db *gorm.DB) {
		db.Order("id").Find(items)
	}
	return Read(selectAll)
}
