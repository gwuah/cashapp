package repository

import "gorm.io/gorm"

type Repo struct {
	Accounts     *accountLayer
	Transactions *transactionLayer
}

func NewRepository(db *gorm.DB) Repo {
	return Repo{
		Accounts:     newAccountLayer(db),
		Transactions: newTransactionLayer(db),
	}
}
