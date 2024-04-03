package database

import (
	"log"
	"performance-dashboard/pkg/database/model"
	"performance-dashboard/pkg/jira/model"
)

func SaveSprint(s *jira.Sprint) error {
	db, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return err
	}
	db.Save(&database.Sprint{ID: s.ID, Name: s.Name, Goal: s.Goal, StartDate: s.StartDate, EndDate: s.EndDate})
	// Create or refress sprintId
	activeSprint := &database.ActiveSprint{SprintID: s.ID}
	db.Where(&database.ActiveSprint{ID: 1}).Assign(activeSprint).FirstOrCreate(activeSprint)
	return nil
}