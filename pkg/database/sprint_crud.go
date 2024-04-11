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
	sprint := database.Sprint{
		ID: s.ID,
		Name: s.Name, 
		Goal: s.Goal, 
		StartDate: s.StartDate, 
		EndDate: s.EndDate,
		State: s.State,
	}
	db.Where(&database.Sprint{ID: s.ID}).Assign(&sprint).FirstOrCreate(&sprint)	// Create or refress sprintId
	return nil
}