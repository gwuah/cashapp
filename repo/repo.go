package repo

import "gorm.io/gorm"

type Repo struct {
	Accounts *accountLayer
}

func NewRepo(db *gorm.DB) Repo {
	return Repo{
		Accounts: newAccountLayer(db),
	}
}
