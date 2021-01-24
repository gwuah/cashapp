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
type Purpose string

var (
	Debit  Type = "debit"
	Credit Type = "credit"

	Failed  Status = "failed"
	Pending Status = "pending"
	Success Status = "success"

	Incoming Direction = "incoming"
	Outgoing Direction = "outgoing"

	Transfer  Purpose = "transfer"
	Deposit   Purpose = "deposit"
	Withrawal Purpose = "withdrawal"
	Reversal  Purpose = "reversal"
)

type Model struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

type User struct {
	Model
	Tag     string   `json:"tag"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Model
	UserID    int   `json:"user_id"`
	User      *User `json:"user,omitempty"`
	IsPrimary bool  `json:"is_primary,omitempty"`
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
	Amount            int64              `json:"amount"`
	Purpose           Purpose            `json:"purpose"`
	TransactionEvents []TransactionEvent `json:"transaction_events"`
}

type TransactionEvent struct {
	Model
	TransactionID int   `json:"transaction_id"`
	WalletID      int   `json:"wallet_id"`
	Type          Type  `json:"type"`
	Amount        int64 `json:"amount"`
}

func RunSeeds(db *gorm.DB) {
	user := User{
		Tag: "yaw",
	}

	if err := db.Model(&User{}).Where("tag=?", user.Tag).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&user)
		} else {
			log.Println("err is nil", err)
		}
	}

	wallet := Wallet{
		UserID:    user.ID,
		IsPrimary: true,
	}

	if err := db.Model(&Wallet{}).Where("account_id=? AND is_primary=?", user.ID, true).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&wallet)
		} else {
			log.Println("err is nil", err)
		}
	}

}
