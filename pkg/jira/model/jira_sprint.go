package jira

import "time"

type Sprint struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	CreatedDate   time.Time `json:"createdDate"`
	ActivatedDate time.Time `json:"activatedDate"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	OriginBoardID int       `json:"originBoardId"`
	Self          string    `json:"self"`
	State         string    `json:"state"`
	Goal          string    `json:"goal"`
}

