package database

import "time"

type Poll struct {
	ID               int `gorm:"unique;primaryKey"`
	Timestamp        time.Time
	ActiveSprint     int
	HeadIssueStateID int
}
