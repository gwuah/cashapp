package repository

import (
	"cashapp/models"

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
