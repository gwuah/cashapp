package services

import (
	"cashapp/core"
	"cashapp/repository"
	"errors"
	"log"
	"net/http"

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

	if err := c.repository.Accounts.Create(account); err != nil {
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
