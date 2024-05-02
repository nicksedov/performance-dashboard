package database

import (
	database "performance-dashboard/pkg/database/model"
	"time"
)

func NewPoll(activeSprint int) (*database.Poll, error) {
	poll := &database.Poll{Timestamp: time.Now(), ActiveSprint: activeSprint, Committed: false}
	db.Save(poll)
	return poll, nil
}

func CommitPoll(poll *database.Poll) error {
	poll.Committed = true
	db.Save(poll)
	return nil
}
