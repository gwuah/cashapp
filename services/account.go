package services

import (
	"cashapp/core"
	"cashapp/repository"
	"errors"

	"gorm.io/gorm"
)

type accountLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newAccountLayer(r repository.Repo, c *core.Config) *accountLayer {
	return &accountLayer{
		repository: r,
		config:     c,
	}
}

func (c *accountLayer) CreateAccount(req core.CreateAccountRequest) core.Response {
	account, err := c.repository.Accounts.FindByTag(req.Tag)

	if err == nil {
		return core.Error(err, core.String("cash tag has already been taken"))
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return core.Error(err, nil)
	}

	if err := c.repository.Accounts.Create(account); err != nil {
		return core.Error(err, nil)

	}

	return core.Success(&map[string]interface{}{
		"account": account,
	}, core.String("account created successfully"))

}
