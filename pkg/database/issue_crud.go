package database

import (
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
	"time"
)

const ISO8601_LAYOUT string = "2006-01-02T15:04:05Z0700"

func SaveIssue(pollId int, iss *jira.Issue, f *jira.IssueFields, parentId int) (*database.IssueState, error) {
	_, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return nil, err
	}
	created, _ := time.Parse(ISO8601_LAYOUT, f.Created)
	actualStart, _ := time.Parse(ISO8601_LAYOUT, f.ActualStart)
	actualEnd, _ := time.Parse(ISO8601_LAYOUT, f.ActualEnd)

	newIssue := database.Issue{
		Key:            iss.Key,
		Type:           f.Issuetype.ID,
		Summary:        f.Summary,
		CreatorID:      f.Creator.AccountID,
		Created:        created,
		ReporterID:     f.Reporter.AccountID,
		Description:    f.Description,
		ActualStart:    actualStart,
		ActualEnd:      actualEnd,
		ActualSprintID: f.Sprint.ID,
		Subtask: 		f.Issuetype.Subtask,
		ParentID:       parentId,
	}

	db.Where(database.Issue{Key: iss.Key}).Assign(newIssue).FirstOrCreate(&newIssue)

	issueStateRecord := &database.IssueState{
		PollID:         pollId,
		IssueID:        newIssue.ID,
		AssigneeID:     f.Assignee.AccountID,
		StoryPoints:    f.StoryPoints,
		StatusCategory: f.Status.StatusCategory.Key,
		StatusID:       f.Status.ID,
	}

	db.Save(issueStateRecord)

	for _, sprint := range f.ClosedSprints {
		dbClosed := database.IssueClosedSprint{
			IssueID:  newIssue.ID,
			SprintID: sprint.ID,
		}
		db.Where(dbClosed).Assign(dbClosed).FirstOrCreate(&dbClosed)
	}
	return issueStateRecord, nil
}
