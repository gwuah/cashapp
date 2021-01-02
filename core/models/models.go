package models

import (
	"errors"
	"log"
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
	Tag     string   `json:"tag"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Model
	AccountID    int           `json:"account_id"`
	Account      *Account      `json:"account,omitempty"`
	IsPrimary    bool          `json:"is_primary,omitempty"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Model
	Amount      int       `json:"amount"`
	AccountID   int       `json:"account_id"`
	Account     *Account  `json:"account,omitempty"`
	WalletID    int       `json:"wallet_id"`
	Wallet      *Wallet   `json:"wallet,omitempty"`
	Description string    `json:"description"`
	Direction   Direction `json:"direction"`
	Status      Status    `json:"status"`
	Date        time.Time `json:"transaction_date"`
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
