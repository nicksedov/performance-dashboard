package database

import "time"

type Issue struct {
	ID             int `gorm:"unique;primaryKey"`
	Key            string
	Type           string
	Summary        string
	CreatorID      string
	Created        time.Time
	ReporterID     string
	Description    string
	ActualStart    time.Time
	ActualEnd      time.Time
	ActualSprintID int
	Subtask        bool `gorm:"default:false"`
	ParentID       int
}

type IssueState struct {
	ID             int `gorm:"unique;primaryKey"`
	PollID         int
	IssueID        int
	AssigneeID     string
	StoryPoints    float64
	StatusCategory string
	StatusID       string
}

type IssueClosedSprint struct {
	ID       int `gorm:"unique;primaryKey"`
	IssueID  int
	SprintID int
}
