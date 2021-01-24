package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	Accounts          *accountLayer
	Transactions      *transactionLayer
	TransactionEvents *eventLayer

	Wallets *walletLayer
}

func NewRepository(db *gorm.DB) Repo {
	return Repo{
		Accounts:          newAccountLayer(db),
		Transactions:      newTransactionLayer(db),
		Wallets:           newWalletLayer(db),
		TransactionEvents: newEventLayer(db),
	}

}
