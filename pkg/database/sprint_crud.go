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
	existing := database.Sprint{}
	tx := db.Where(database.Sprint{ID: sprint.ID}).First(&existing)
	if tx.Error == nil {
		sprint.ID = existing.ID
		if !existing.Equals(&sprint) {
			db.Save(&sprint)
		} else {
			log.Printf("Sprint with ID '%d' is already known\n", sprint.ID)
		}
	} else {
		db.Save(&sprint)
	}

	return nil
}