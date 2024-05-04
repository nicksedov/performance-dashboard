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
	//LastPollID    int        `gorm:"default:0"`
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
		// must not compare by LastPollID 
}

type SprintPoll struct {
	ID              int        `gorm:"unique;primaryKey"`
	LastPollID      int
	CompletionPoll  bool       `gorm:"default:false"`
}
