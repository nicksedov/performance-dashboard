package dto

type IssueType struct {
	ID               int   `gorm:"unique;primaryKey"`
	Type           string  
	TypeName       string  
}