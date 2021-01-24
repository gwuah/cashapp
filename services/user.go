package services

import (
	"cashapp/core"
	"cashapp/repository"
	"errors"

	"gorm.io/gorm"
)

type userLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newUserLayer(r repository.Repo, c *core.Config) *userLayer {
	return &userLayer{
		repository: r,
		config:     c,
	}
}

func (c *userLayer) CreateUser(req core.CreateUserRequest) core.Response {
	user, err := c.repository.Users.FindByTag(req.Tag)

	if err == nil {
		return core.Error(err, core.String("cash tag has already been taken"))
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return core.Error(err, nil)
	}

	if err := c.repository.Users.Create(user); err != nil {
		return core.Error(err, nil)
	}

	// an automatic wallet is created for a every new user
	wallet, err := c.repository.Wallets.Create(user.ID)
	if err != nil {
		return core.Error(err, nil)

	}

	user.Wallets = append(user.Wallets, *wallet)

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("user created successfully"))

}
