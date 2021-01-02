package repo

import (
	"cashapp/core/models"

	"gorm.io/gorm"
)

type accountLayer struct {
	db *gorm.DB
}

func newAccountLayer(db *gorm.DB) *accountLayer {
	return &accountLayer{
		db: db,
	}
}

func (al *accountLayer) Create(account *models.Account) error {
	if err := al.db.Create(account).Error; err != nil {
		return err
	}
	return nil

}

func (al *accountLayer) FindByTag(tag string) (*models.Account, error) {
	account := models.Account{Tag: tag}
	if err := al.db.Where("tag = ?", tag).First(&account).Error; err != nil {
		return &account, err
	}
	return &account, nil
}
