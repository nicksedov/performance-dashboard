package database

import (
	"log"
	"time"
	"slices"
	database "performance-dashboard/pkg/database/model"
	jira "performance-dashboard/pkg/jira/model"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type WhereClause func() database.Account

const ISO8601_LAYOUT string = "2006-01-02T15:04:05Z0700"

var accountIdCache *cache.Cache = cache.New(30*time.Second, 1*time.Minute)
var accountNameCache *cache.Cache = cache.New(30*time.Second, 1*time.Minute)

func SaveIssue(pollId int, iss *jira.Issue, f *jira.IssueFields, parentId int) *database.IssueState {
	
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
		Subtask:        f.Issuetype.Subtask,
		ParentID:       parentId,
		CurrentState:   f.Status.StatusCategory.Key,
	}

	existing := database.Issue{}
	tx := db.Where(database.Issue{Key: iss.Key}).First(&existing)
	if tx.Error == nil {
		newIssue.ID = existing.ID
		if !existing.Equals(&newIssue) {
			db.Save(&newIssue)
		} else {
			log.Printf("Issue with key '%s' is already known\n", iss.Key)
		}
	} else {
		db.Save(&newIssue)
	}

	issueStateRecord := saveIssueState(pollId, newIssue.ID, assignee.ID, f)
	saveIssueSprints(newIssue.ID, f)
	saveOrUpdateAssigneeTransitions(newIssue.ID, assignee.ID)
	return issueStateRecord
}

func getAccountMetadata(acc *jira.Account) *database.Account {
	var result *database.Account = &database.Account{}
	if acc.AccountID != "" {
		result = getCached(accountIdCache, acc.AccountID, func() database.Account { return database.Account{AccountID: acc.AccountID} })
	} else if acc.DisplayName != "" {
		result = getCached(accountNameCache, acc.DisplayName, func() database.Account { return database.Account{DisplayName: acc.DisplayName} })
	}
	return result
}

func getCached(accountCache *cache.Cache, key string, whereClause WhereClause) *database.Account {
	var result database.Account
	resultObj, found := accountCache.Get(key)
	if found {
		result = resultObj.(database.Account)
	} else {
		r := db.Where(whereClause()).First(&result)
		if r.Error == nil {
			accountNameCache.Add(key, result, cache.DefaultExpiration)
		}
	}
	return &result
}

func saveIssueState(pollId int, issueId int, assigneeId int, f *jira.IssueFields) *database.IssueState {
	issueStateRecord := database.IssueState{
		PollID:         pollId,
		IssueID:        issueId,
		AssigneeID:     assigneeId,
		StoryPoints:    f.StoryPoints,
		StatusCategory: f.Status.StatusCategory.Key,
		StatusID:       f.Status.ID,
	}
	db.Save(&issueStateRecord)
	return &issueStateRecord
}

func saveIssueSprints(issueId int, f *jira.IssueFields) {
	existingIssueSprints,_ := Read(
		func(items *[]database.IssueSprint, db *gorm.DB) {
			db.Where(database.IssueSprint{IssueID: issueId}).Find(items)
		})
	existingSprintIds := make([]int, 0, len(*existingIssueSprints))
	for _, existingSprint := range *existingIssueSprints {
		existingSprintIds = append(existingSprintIds, existingSprint.SprintID)
	}
	// Closed sprints
	for _, sprint := range f.ClosedSprints {
		spID := sprint.ID
		if spID != 0 && !slices.Contains(existingSprintIds, spID) {
			db.Save(&database.IssueSprint{IssueID: issueId, SprintID: spID })
		}
	}
	// Current sprint
	spID := f.Sprint.ID
	if spID != 0 && !slices.Contains(existingSprintIds, spID) {
		db.Save(&database.IssueSprint{IssueID: issueId, SprintID: spID })
	}
}

func saveOrUpdateAssigneeTransitions(issueId int, assigneeId int) {
	
	if assigneeId == 0 {
		return; // Assignee absence or removal is not a transition 
	}

	existingTransitions,_ := Read(
		func(items *[]database.IssueAssigneeTransitions, db *gorm.DB) {
			db.Where(database.IssueAssigneeTransitions{IssueID: issueId}).Find(items)
		})
	transitionsExist := len(*existingTransitions) > 0
	if transitionsExist {
		update := (*existingTransitions)[0]
		if update.LastAssigneeID != assigneeId {
			update.Transitions = update.Transitions + 1
			update.LastAssigneeID = assigneeId
			db.Save(&update)
		} else {
			newTransition := database.IssueAssigneeTransitions{IssueID: issueId, LastAssigneeID: assigneeId}
			db.Save(&newTransition)
		}
	}
}
