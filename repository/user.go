package repository

import (
	"cashapp/models"

	"gorm.io/gorm"
)

type userLayer struct {
	db *gorm.DB
}

func newUserLayer(db *gorm.DB) *userLayer {
	return &userLayer{
		db: db,
	}
}

func (ul *userLayer) Create(user *models.User) error {
	if err := ul.db.Create(user).Error; err != nil {
		return err
	}
	return nil

}

func (ul *userLayer) FindByTag(tag string) (*models.User, error) {
	user := models.User{Tag: tag}
	if err := ul.db.Where("tag = ?", tag).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}
