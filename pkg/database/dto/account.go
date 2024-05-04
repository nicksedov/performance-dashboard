package dto

type Account struct {
	ID           int     `gorm:"unique;primaryKey"`
	AccountID    string
	AccountType  string
	Role         string     
	DisplayName  string  `gorm:"index:idx_account_displayname"`
	EmailAddress string     
}