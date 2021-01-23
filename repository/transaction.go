package repository

import (
	"cashapp/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type transactionLayer struct {
	db *gorm.DB
}

func newTransactionLayer(db *gorm.DB) *transactionLayer {
	return &transactionLayer{
		db: db,
	}
}

func (tl *transactionLayer) SQLTransaction(f func(tx *gorm.DB) error) error {
	return tl.db.Transaction(f)
}

func (tl *transactionLayer) Create(data *models.Transaction) error {
	if err := tl.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (tl *transactionLayer) GetWalletBalance(id int) (int64, error) {
	var balance int64

	rows, err := tl.db.Table("transaction_events").Select("amount, type").Where("wallet_id = ?", id).Rows()
	if err != nil {
		return balance, err
	}
	defer rows.Close()

	for rows.Next() {
		var amount int64
		var event_type string
		if err := rows.Scan(&amount, &event_type); err != nil {
			return 0, fmt.Errorf("error reading amount/type: %v", err)
		}
		if strings.EqualFold(event_type, string(models.Debit)) {
			balance -= amount
		} else {
			balance += amount
		}
	}

	return balance, nil
}
