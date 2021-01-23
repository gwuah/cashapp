package models

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type Type string
type Status string
type Direction string

var (
	Debit  Type = "debit"
	Credit Type = "credit"

	Failed  Status = "failed"
	Pending Status = "pending"
	Success Status = "success"

	Incoming Direction = "incoming"
	Outgoing Direction = "outgoing"
)

type Model struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

type Account struct {
	Model
	Tag     string   `json:"tag"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Model
	AccountID int      `json:"account_id"`
	Account   *Account `json:"account,omitempty"`
	IsPrimary bool     `json:"is_primary,omitempty"`
}

type Transaction struct {
	Model
	Direction         Direction          `json:"direction"`
	Status            Status             `json:"status"`
	Description       string             `json:"description"`
	Ref               string             `json:"ref"`
	From              int                `json:"from"`
	To                int                `json:"to"`
	WalletID          int                `json:"wallet_id"`
	AccountID         int                `json:"account_id"`
	Account           *Account           `json:"account,omitempty"`
	TransactionEvents []TransactionEvent `json:"transaction_lines"`
	Amount            int64              `json:"amount"`
}

type TransactionEvent struct {
	Model
	TransactionID string `json:"transaction_id"`
	WalletID      string `json:"wallet_id"`
	Type          Type   `json:"type"`
	Amount        int64  `json:"amount"`
}

func RunSeeds(db *gorm.DB) {
	account := Account{
		Tag: "yaw",
	}

	if err := db.Model(&Account{}).Where("tag=?", account.Tag).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&account)
		} else {
			log.Println("err is nil", err)
		}
	}

	wallet := Wallet{
		AccountID: account.ID,
		IsPrimary: true,
	}

	if err := db.Model(&Wallet{}).Where("account_id=? AND is_primary=?", account.ID, true).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&wallet)
		} else {
			log.Println("err is nil", err)
		}
	}

}
