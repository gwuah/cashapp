package repository

import (
	"cashapp/models"

	"gorm.io/gorm"
)

type walletLayer struct {
	db *gorm.DB
}

func newWalletLayer(db *gorm.DB) *walletLayer {
	return &walletLayer{
		db: db,
	}
}

func (wl *walletLayer) Create(userId int) (*models.Wallet, error) {
	wallet := models.Wallet{
		UserID:    userId,
		IsPrimary: true,
	}

	if err := wl.db.Create(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (wl *walletLayer) FindPrimaryWallet(userId int) (*models.Wallet, error) {
	wallet := models.Wallet{
		UserID:    userId,
		IsPrimary: true,
	}
	if err := wl.db.First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
