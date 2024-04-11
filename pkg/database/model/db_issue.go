package database

import "time"

type Issue struct {
	ID            int        `gorm:"unique;primaryKey"`
	Key           string
	CreatorID     string
	CreatedDate   string     `gorm:"type:date"`
	ReporterID    string
	Description   string
	ActualStart   time.Time
	ActualEnd     time.Time
}

type IssueHistory struct {
	ID             int        `gorm:"unique;primaryKey"`
	Timestamp      time.Time
	IssueID        int
	AssigneeID     string
	StoryPoints    float64
	StatusCategory string
	StatusID       string 
}

type IssueStatus struct {
	ID   string
	Name string
}
