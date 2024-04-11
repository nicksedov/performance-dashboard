package database

import (
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
	"time"
)

const ISO8601_LAYOUT string = "2006-01-02T15:04:05Z0700"

func SaveIssue(iss *jira.Issue, f *jira.IssueFields) error {
	_, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return err
	}
	actualStart,_ := time.Parse(ISO8601_LAYOUT, f.ActualStart)
	actualEnd,_ := time.Parse(ISO8601_LAYOUT, f.ActualEnd)

	SaveAccount(&f.Creator)
	SaveAccount(&f.Reporter)
	SaveAccount(&f.Assignee)

	newIssue := &database.Issue{
		Key: iss.Key,
		CreatorID:   f.Creator.AccountID,
		CreatedDate: f.Created,
		ReporterID : f.Reporter.AccountID,
		Description: f.Description,
		ActualStart: actualStart,
		ActualEnd  : actualEnd,
	}

	db.Where(&database.Issue{Key: iss.Key}).Assign(newIssue).FirstOrCreate(newIssue)

	issueCurrentState := &database.IssueHistory{
		Timestamp:   time.Now(),
		IssueID:     newIssue.ID,
		AssigneeID:  f.Assignee.AccountID,
		StoryPoints: f.StoryPoints,
		StatusCategory: f.Status.StatusCategory.Key,
		StatusID: f.Status.ID,
	}

	db.Save(issueCurrentState)
	return nil
}