package services

import (
	"cashapp/infra"
	"cashapp/repo"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type accountLayer struct {
	repo   repo.Repo
	config *infra.Config
}

func newAccountLayer(r repo.Repo, c *infra.Config) *accountLayer {
	return &accountLayer{
		repo:   r,
		config: c,
	}
}

func (c *accountLayer) CreateAccount(req infra.CreateAccountRequest) infra.Response {
	account, err := c.repo.Accounts.FindByTag(req.Tag)

	if err == nil {
		return infra.Response{
			Error: true,
			Code:  http.StatusBadRequest,
			Meta: infra.Meta{
				Data:    nil,
				Message: "cash tag has already been taken.",
			},
		}
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("findByTag failed. err", err)
		return infra.Response{
			Error: true,
			Code:  http.StatusInternalServerError,
			Meta: infra.Meta{
				Data:    nil,
				Message: "request failed",
			},
		}
	}

	err = c.repo.Accounts.Create(account)
	if err != nil {
		log.Println("create account failed. err", err)
		return infra.Response{
			Error: true,
			Code:  http.StatusInternalServerError,
			Meta: infra.Meta{
				Data:    nil,
				Message: "request failed",
			},
		}
	}

	return infra.Response{
		Error: false,
		Code:  http.StatusOK,
		Meta: infra.Meta{
			Data: map[string]interface{}{
				"account": account,
			},
			Message: "account created successfully",
		},
	}
}
