package model

import "time"

type Sprint struct {
	ID            int       `json:"id"`
	EndDate       time.Time `json:"endDate"`
	OriginBoardID int       `json:"originBoardId"`
	Self          string    `json:"self"`
	State         string    `json:"state"`
	Name          string    `json:"name"`
	StartDate     time.Time `json:"startDate"`
	CreatedDate   time.Time `json:"createdDate"`
	Goal          string    `json:"goal"`
}