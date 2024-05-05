package database

import (
	"log"
	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/jira/model"
)

func SaveSprint(s *model.Sprint) {

	sprint := dto.Sprint{
		ID:          s.ID,
		Name:        s.Name,
		Goal:        s.Goal,
		CreatedDate: s.CreatedDate,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
		State:       s.State,
	}
	existing := dto.Sprint{}
	tx := db.Where(dto.Sprint{ID: s.ID}).First(&existing)
	if tx.Error == nil {
		if !existing.Equals(&sprint) {
			db.Save(&sprint)
		} else {
			log.Printf("Sprint with ID '%d' is already known\n", sprint.ID)
		}
	} else {
		log.Printf("A new sprint with ID '%d' and name '%s' will be created\n", s.ID, sprint.Name)
		db.Save(&sprint)
	}
}