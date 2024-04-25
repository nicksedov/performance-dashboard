package database

import (
	"log"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
	"github.com/patrickmn/go-cache"
	"time"
)

const ISO8601_LAYOUT string = "2006-01-02T15:04:05Z0700"

var accountIdCache cache.Cache = *cache.New(30*time.Second, 1*time.Minute)
var accountNameCache cache.Cache = *cache.New(30*time.Second, 1*time.Minute)

func SaveIssue(pollId int, iss *jira.Issue, f *jira.IssueFields, parentId int) (*database.IssueState, error) {
	_, err := initDb()
	if err != nil {
		log.Println("Warning: failed to connect database")
		return nil, err
	}
	created, _ := time.Parse(ISO8601_LAYOUT, f.Created)
	actualStart, _ := time.Parse(ISO8601_LAYOUT, f.ActualStart)
	actualEnd, _ := time.Parse(ISO8601_LAYOUT, f.ActualEnd)

	creator := getAccountMetadata(&f.Creator)
	reporter := getAccountMetadata(&f.Reporter)
	assignee := getAccountMetadata(&f.Assignee)

	newIssue := database.Issue{
		Key:            iss.Key,
		Type:           f.Issuetype.ID,
		Summary:        f.Summary,
		CreatorID:      creator.ID,
		Created:        created,
		ReporterID:     reporter.ID,
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
		AssigneeID:     assignee.ID,
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

func getAccountMetadata(acc *jira.Account) *database.Account {
	var result database.Account
	if acc.AccountID != "" {
		resultObj, found := accountIdCache.Get(acc.AccountID)
		if found {
			result = resultObj.(database.Account)
		} else {
			r := db.Where("account_id = ?", acc.AccountID).First(&result)
			if r.Error == nil {
				accountIdCache.Add(acc.AccountID, result, cache.DefaultExpiration)
			}
		}
	} else if acc.DisplayName != "" {
		resultObj, found := accountNameCache.Get(acc.DisplayName)
		if found {
			result = resultObj.(database.Account)
		} else {
			r := db.Where("display_name = ?", acc.DisplayName).First(&result)
			if r.Error == nil {
				accountNameCache.Add(acc.DisplayName, result, cache.DefaultExpiration)
			}
		}
	}
	return &result
}