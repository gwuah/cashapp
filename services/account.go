package services

import (
	"cashapp/core"
	"cashapp/repo"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type accountLayer struct {
	repo   repo.Repo
	config *core.Config
}

func newAccountLayer(r repo.Repo, c *core.Config) *accountLayer {
	return &accountLayer{
		repo:   r,
		config: c,
	}
}

func (c *accountLayer) CreateAccount(req core.CreateAccountRequest) core.Response {
	account, err := c.repo.Accounts.FindByTag(req.Tag)

	if err == nil {
		return core.Response{
			Error: true,
			Code:  http.StatusBadRequest,
			Meta: core.Meta{
				Data:    nil,
				Message: "cash tag has already been taken.",
			},
		}
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("findByTag failed. err", err)
		return core.Response{
			Error: true,
			Code:  http.StatusInternalServerError,
			Meta: core.Meta{
				Data:    nil,
				Message: "request failed",
			},
		}
	}

	err = c.repo.Accounts.Create(account)
	if err != nil {
		log.Println("create account failed. err", err)
		return core.Response{
			Error: true,
			Code:  http.StatusInternalServerError,
			Meta: core.Meta{
				Data:    nil,
				Message: "request failed",
			},
		}
	}

	return core.Response{
		Error: false,
		Code:  http.StatusOK,
		Meta: core.Meta{
			Data: map[string]interface{}{
				"account": account,
			},
			Message: "account created successfully",
		},
	}
}
