package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	Users             *userLayer
	Transactions      *transactionLayer
	TransactionEvents *eventLayer
	Wallets           *walletLayer
}

func NewRepository(db *gorm.DB) Repo {
	return Repo{
		Users:             newUserLayer(db),
		Transactions:      newTransactionLayer(db),
		Wallets:           newWalletLayer(db),
		TransactionEvents: newEventLayer(db),
	}

}
