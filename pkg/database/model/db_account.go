package database

type Account struct {
	ID           string     `gorm:"unique;primaryKey"`
	AccountType  string
	Role         string     
	DisplayName  string     
}