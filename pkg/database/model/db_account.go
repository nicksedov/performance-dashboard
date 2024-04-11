package database

type Account struct {
	ID           string     `gorm:"unique;primaryKey"`
	AccountType  string     
	Active       bool       
	DisplayName  string     
	EmailAddress string     
	Key          string     
	Name         string     
	TimeZone     string     
}