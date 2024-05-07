package dto

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
	LastSprintID   int
	Subtask        bool `gorm:"default:false"`
	ParentID       int
	EpicID         int
	CurrentState   string
}

func (it *Issue) Equals(that *Issue) bool {
	return it.Key == that.Key &&
		it.Type == that.Type &&
		it.Summary == that.Summary &&
		it.CreatorID == that.CreatorID &&
		it.Created.Equal(that.Created) &&
		it.ReporterID == that.ReporterID &&
		it.Description == that.Description &&
		it.ActualStart.Equal(that.ActualStart) &&
		it.ActualEnd.Equal(that.ActualEnd) &&
		it.LastSprintID == that.LastSprintID &&
		it.Subtask == that.Subtask &&
		it.ParentID == that.ParentID &&
		it.EpicID == that.EpicID &&
		it.CurrentState == that.CurrentState
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

type IssueSprint struct {
	ID       int `gorm:"unique;primaryKey"`
	IssueID  int
	SprintID int
}

type IssueAssigneeTransitions struct {
	IssueID        int `gorm:"unique;primaryKey"`
	LastAssigneeID int
	Transitions    int `gorm:"default:0"`
}
