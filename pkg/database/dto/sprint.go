package dto

import (
	"time"
)

type Sprint struct {
	ID            int        `gorm:"unique;primaryKey"`
	Name          string
	Goal          string
	CreatedDate   time.Time
	ActivatedDate time.Time
	StartDate     time.Time
	EndDate       time.Time
	State         string
}

func (it *Sprint) Equals(that *Sprint) bool {
	return it.ID == that.ID && 
		it.Name == that.Name && 
		it.Goal == that.Goal && 
		it.CreatedDate.Equal(that.CreatedDate) &&
		it.ActivatedDate.Equal(that.ActivatedDate) &&
		it.StartDate.Equal(that.StartDate) &&
		it.EndDate.Equal(that.EndDate) &&
		it.State == that.State
}

type SprintPoll struct {
	ID              int        `gorm:"unique;primaryKey"`
	FirstPollID     int
	LastPollID      int
	CompletionPoll  bool       `gorm:"default:false"`
}
