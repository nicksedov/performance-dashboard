package database

import (
	"log"
	database "performance-dashboard/pkg/database/model"
	"time"
)

func NewPoll(activeSprint int) (*database.Poll, error) {
	db, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return nil, err
	}
	poll := &database.Poll{Timestamp: time.Now(), ActiveSprint: activeSprint, Committed: false}
	db.Save(poll)
	return poll, nil
}

func CommitPoll(poll *database.Poll) error {
	db, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return err
	}
	poll.Committed = true
	db.Save(poll)
	return nil
}
