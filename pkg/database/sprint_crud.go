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

func CompletionPollRequired(sprintID int) bool {
	sprintPoll := dto.SprintPoll{CompletionPoll: false}
	db.Where(dto.SprintPoll{ID: sprintID}).First(&sprintPoll)
	return !sprintPoll.CompletionPoll
}

func UpdateSprintPoll(sprintID int, pollId int, isCompletionPoll bool) {
	sprintPoll := dto.SprintPoll{
		ID:             sprintID,
		LastPollID:     pollId,
		CompletionPoll: isCompletionPoll,
	}
	db.Save(&sprintPoll)
}
