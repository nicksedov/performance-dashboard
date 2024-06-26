package database

import (
	"performance-dashboard/pkg/database/dto"
	"time"
)

func NewPoll(activeSprint int) (*dto.Poll, error) {
	poll := &dto.Poll{Timestamp: time.Now(), ActiveSprint: activeSprint, Committed: false}
	GetDB().Save(poll)
	return poll, nil
}

func CommitPoll(poll *dto.Poll) error {
	poll.Committed = true
	GetDB().Save(poll)
	return nil
}
