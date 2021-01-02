package models

import (
	"time"

	"gorm.io/gorm"
)

type Direction string
type Status string

var (
	Debit  Direction = "debit"
	Credit Direction = "credit"

	Failed  Status = "failed"
	Success Status = "success"
)

type Model struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

type Account struct {
	Model
	Tag          string        `json:"tag"`
	Balance      int64         `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Model
	Amount      int       `json:"amount"`
	AccountID   int       `json:"account_id"`
	Account     *Account  `json:"account,omitempty"`
	Description string    `json:"description"`
	Direction   Direction `json:"direction"`
	Status      Status    `json:"status"`
	Date        time.Time `json:"transaction_date"`
}

func RunSeeds(db *gorm.DB) {

	db.Create(&Account{
		Tag:     "yaw",
		Balance: 0,
	})

}
