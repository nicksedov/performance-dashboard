package database

import "time"

type ApplicationLog struct {
	ID           int     `gorm:"unique;primaryKey"`
	Timestamp    time.Time
	Log          string
}