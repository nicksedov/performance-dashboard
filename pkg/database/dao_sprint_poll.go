package database

import "performance-dashboard/pkg/database/dto"

func SaveSprintPoll(sprintID int, pollId int, isCompletionPoll bool) {
	sprintPoll := dto.SprintPoll{
		ID:             sprintID,
		LastPollID:     pollId,
		CompletionPoll: isCompletionPoll,
	}
	db.Save(&sprintPoll)
}

// Check that a final poll was made for a given sprint after its completion  
func CompletionPollRequired(sprintID int) bool {
	sprintPoll := dto.SprintPoll{CompletionPoll: false}
	db.Where(dto.SprintPoll{ID: sprintID}).First(&sprintPoll)
	return !sprintPoll.CompletionPoll
}
