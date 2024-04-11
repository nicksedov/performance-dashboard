package database

import (
	"time"
)

type Sprint struct {
	ID            int        `gorm:"unique;primaryKey"`
	Name          string
	Goal          string
	CreatedDate   time.Time
	ActivatedDate time.Time
	StartDate     time.Time
	EndDate       time.Time
	State         string
}

 