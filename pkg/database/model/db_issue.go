package database

import "time"

type Issue struct {
	ID             int `gorm:"unique;primaryKey"`
	Key            string
	Type           string
	Summary        string
	CreatorID      int
	Created        time.Time
	ReporterID     int
	Description    string
	ActualStart    time.Time
	ActualEnd      time.Time
	ActualSprintID int
	Subtask        bool `gorm:"default:false"`
	ParentID       int
}

func (this *Issue) Equals(that *Issue) bool {
	return this.Key == that.Key && 
		this.Type == that.Type && 
		this.Summary == that.Summary &&
		this.CreatorID == that.CreatorID &&
		this.Created.Equal(that.Created) &&
		this.ReporterID == that.ReporterID &&
		this.Description == that.Description &&
		this.ActualStart.Equal(that.ActualStart) &&
		this.ActualEnd.Equal(that.ActualEnd) &&
		this.ActualSprintID == that.ActualSprintID &&
		this.Subtask == that.Subtask &&
		this.ParentID == that.ParentID
}

type IssueState struct {
	ID             int `gorm:"unique;primaryKey"`
	PollID         int
	IssueID        int
	AssigneeID     int
	StoryPoints    float64
	StatusCategory string
	StatusID       string
}

type IssueClosedSprint struct {
	ID       int `gorm:"unique;primaryKey"`
	IssueID  int
	SprintID int
}
