package database

import (
	"performance-dashboard/pkg/database/dto"

	"gorm.io/gorm"
)

func SaveSprintPoll(sprintID int, pollId int, isCompletionPoll bool) {
	sprintPoll := dto.SprintPoll{}
	err := GetDB().Where(dto.SprintPoll{ID: sprintID}).First(&sprintPoll).Error
	if err == gorm.ErrRecordNotFound {
		sprintPoll.FirstPollID = pollId
	}
	sprintPoll.ID = sprintID
	sprintPoll.LastPollID = pollId
	sprintPoll.CompletionPoll = isCompletionPoll
	GetDB().Save(&sprintPoll)
}

// Check that a final poll was made for a given sprint after its completion
func CompletionPollRequired(sprintID int) bool {
	sprintPoll := dto.SprintPoll{}
	tx := GetDB().Where(dto.SprintPoll{ID: sprintID}).First(&sprintPoll)
	if tx.Error == nil {
		return !sprintPoll.CompletionPoll
	} else {
		return false
	}
}
