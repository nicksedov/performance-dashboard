package database

import "time"

type Issue struct {
	ID            int        `gorm:"unique;primaryKey"`
	Key           string
	Goal          string
	CreatedDate   time.Time
}

