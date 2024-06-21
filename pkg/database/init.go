package database

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"time"

	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/profiles"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"github.com/DATA-DOG/go-sqlmock"
)

var mockDb *gorm.DB
var db []*gorm.DB
var activeNode int
var masterNodeDetectionTimestamp = time.Unix(0,0)

func InitializeDB() error {
	dbConfigs := profiles.GetSettings().DbConfig
	if len(dbConfigs.DbNode) == 0 {
		return errors.New("Database connection parameters not configured")
	} 
	db = make([]*gorm.DB, 0, len(dbConfigs.DbNode))
	for _, dbConfig := range dbConfigs.DbNode {
		postgresCfg := buildPostgresConfig(dbConfig)
		dbNode, err := gorm.Open(postgres.New(*postgresCfg), buildGormConfig(dbConfig))
		if err == nil {
			log.Printf("Database connection established to postgres://%s:%d/%s\n",
					dbConfig.Host, dbConfig.Port, dbConfig.DbName)
			db = append(db, dbNode)
			if sqlDb, dbErr := dbNode.DB(); dbErr == nil {
				sqlDb.SetMaxIdleConns(2)
				sqlDb.SetMaxOpenConns(2)
				dbStats, statsErr := json.MarshalIndent(sqlDb.Stats(), "", "  ");  if statsErr == nil {
					log.Printf("Database connection statistics: %s\n", string(dbStats))
				}
			} else {
				err = dbErr
			}
		}
		if err != nil {
			log.Printf("Warning: Database on host %s:%d is unavailable\n", dbConfig.Host, dbConfig.Port)
			return err
		} 
	}
	if discoverErr := detectMasterNode(); discoverErr == nil {
		return autoMigrate()
	} else {
		return discoverErr
	}
}

func CloseDB() {
	for _, dbNode := range db {
		if sqlDB, err := dbNode.DB(); err != nil {
			sqlDB.Close()
		}
	}
	if mockDb != nil {
		if sqlDB, err := mockDb.DB(); err != nil {
			sqlDB.Close()
		}
	}
}

func detectMasterNode() error {
	current := time.Now() 
	if current.After(masterNodeDetectionTimestamp) {
		activeNodes := make([]int, 0, len(db))
		for i, dbNode := range db {
			sqlDb, err := dbNode.DB(); 
			if err == nil {
				row := sqlDb.QueryRow("SELECT pg_is_in_recovery()")
				var recovery bool
				if scanErr := row.Scan(&recovery); scanErr == nil {
					if !recovery {
						activeNodes = append(activeNodes, i)
					} 
				} else {
					err = scanErr
				}
			}
			if err != nil {
				continue
			}
		}
		if len(activeNodes) == 0 {
			return errors.New("Database connection not available")
		} else if !slices.Contains(activeNodes, activeNode) {
			activeNode = activeNodes[0]
			log.Printf("Database node %d assigned as master\n", activeNode)
		}
		period := profiles.GetSettings().DbConfig.MasterDetectionPeriod
		if period <= 0 {
			period = time.Minute * 5
		}
		masterNodeDetectionTimestamp = current.Add(period)
	}
	return nil
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
	err := detectMasterNode(); if err == nil {
		if activeNode > -1 && activeNode < len(db) {
			return db[activeNode]
		}
	}
	return getOrCreateMockDB()
}

func getOrCreateMockDB() *gorm.DB {
	if mockDb == nil {
		sqlDb, _, _ := sqlmock.New()
 		dialector := postgres.New(postgres.Config{
  			Conn:       sqlDb,
  			DriverName: "postgres",
 		})
 		mockDb, _ = gorm.Open(dialector, &gorm.Config{})
	}
	return mockDb
}

func buildPostgresConfig(dbConfig profiles.Database) *postgres.Config {
	dsnFormat := "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s"
	dsn := fmt.Sprintf(dsnFormat,
		dbConfig.Host, dbConfig.Port, dbConfig.DbName,
		dbConfig.User, dbConfig.Password, dbConfig.SSLMode)
	postgresCfg := &postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	return postgresCfg
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
