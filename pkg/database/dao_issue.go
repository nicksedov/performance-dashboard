package database

import (
	"errors"
	"log"
	"performance-dashboard/pkg/database/dto"
	"performance-dashboard/pkg/jira/model"
	"slices"
	"time"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type WhereClause [T any]func() T

const ISO8601_LAYOUT string = "2006-01-02T15:04:05Z0700"

var accountIdCache *cache.Cache = cache.New(30*time.Second, 1*time.Minute)
var accountNameCache *cache.Cache = cache.New(30*time.Second, 1*time.Minute)

func SaveIssue(pollId int, iss *model.Issue, f *model.IssueFields, parentId int) *dto.IssueState {

	created, _ := time.Parse(ISO8601_LAYOUT, f.Created)
	actualStart, _ := time.Parse(ISO8601_LAYOUT, f.ActualStart)
	actualEnd, _ := time.Parse(ISO8601_LAYOUT, f.ActualEnd)

	creator := getAccountMetadata(&f.Creator)
	reporter := getAccountMetadata(&f.Reporter)
	assignee := getAccountMetadata(&f.Assignee)

	newIssue := dto.Issue{
		Key:          iss.Key,
		Type:         f.Issuetype.ID,
		Summary:      f.Summary,
		CreatorID:    creator.ID,
		Created:      created,
		ReporterID:   reporter.ID,
		Description:  f.Description,
		ActualStart:  actualStart,
		ActualEnd:    actualEnd,
		LastSprintID: f.Sprint.ID,
		Subtask:      f.Issuetype.Subtask,
		ParentID:     parentId,
		EpicID:       f.Epic.ID,
		CurrentState: f.Status.StatusCategory.Key,
	}

	existing := dto.Issue{}
	tx := db.Where(dto.Issue{Key: iss.Key}).First(&existing)
	if tx.Error == nil {
		newIssue.ID = existing.ID
		if newIssue.LastSprintID == 0 && existing.LastSprintID != 0 {
			newIssue.LastSprintID = existing.LastSprintID
		}
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

func getAccountMetadata(acc *model.Account) *dto.Account {
	var result *dto.Account = &dto.Account{}
	var err error
	if acc.AccountID != "" {
		whereClause := func() dto.Account { return dto.Account{AccountID: acc.AccountID} }
		result, err = getCached(accountIdCache, acc.AccountID, whereClause)
	} else if acc.DisplayName != "" {
		whereClause := func() dto.Account { return dto.Account{DisplayName: acc.DisplayName} }
		result, err = getCached(accountNameCache, acc.DisplayName, whereClause)
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		SaveExternalParticipantAccount(acc)	
	}
	return result
}

func getCached[T any](memCache *cache.Cache, key string, whereClause WhereClause[T]) (*T, error) {
	var result T
	resultObj, found := memCache.Get(key)
	if found {
		result = resultObj.(T)
	} else {
		r := db.Where(whereClause()).First(&result)
		if r.Error == nil {
			accountNameCache.Add(key, result, cache.DefaultExpiration)
		} else {
			return nil, r.Error
		} 
	}
	return &result, nil
}

func saveIssueState(pollId int, issueId int, assigneeId int, f *model.IssueFields) *dto.IssueState {
	issueStateRecord := dto.IssueState{
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

func saveIssueSprints(issueId int, f *model.IssueFields) {
	existingIssueSprints, _ := Read(
		func(items *[]dto.IssueSprint, db *gorm.DB) {
			db.Where(dto.IssueSprint{IssueID: issueId}).Find(items)
		})
	existingSprintIds := make([]int, 0, len(*existingIssueSprints))
	for _, existingSprint := range *existingIssueSprints {
		existingSprintIds = append(existingSprintIds, existingSprint.SprintID)
	}
	// Closed sprints
	for _, sprint := range f.ClosedSprints {
		spID := sprint.ID
		if spID != 0 && !slices.Contains(existingSprintIds, spID) {
			db.Save(&dto.IssueSprint{IssueID: issueId, SprintID: spID})
		}
	}
	// Current sprint
	spID := f.Sprint.ID
	if spID != 0 && !slices.Contains(existingSprintIds, spID) {
		db.Save(&dto.IssueSprint{IssueID: issueId, SprintID: spID})
	}
}

func saveOrUpdateAssigneeTransitions(issueId int, assigneeId int) {

	if assigneeId == 0 {
		return // Assignee absence or removal is not a transition
	}

	existingTransitions, _ := Read(
		func(items *[]dto.IssueAssigneeTransitions, db *gorm.DB) {
			db.Where(dto.IssueAssigneeTransitions{IssueID: issueId}).Find(items)
		})
	transitionsExist := len(*existingTransitions) > 0
	if transitionsExist {
		update := (*existingTransitions)[0]
		if update.LastAssigneeID != assigneeId {
			update.Transitions = update.Transitions + 1
			update.LastAssigneeID = assigneeId
			db.Save(&update)
		}
	} else {
		newTransition := dto.IssueAssigneeTransitions{IssueID: issueId, LastAssigneeID: assigneeId}
		db.Save(&newTransition)
	}
}
