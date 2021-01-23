package services

import (
	"cashapp/core"
	"cashapp/repository"

	"github.com/go-redis/redis/v8"
)

type Services struct {
	Accounts *accountLayer
	Payments *paymentLayer
}

func NewService(r repository.Repo, kvStore *redis.Client, c *core.Config) Services {
	return Services{
		Accounts: newAccountLayer(r, c),
		Payments: newPaymentLayer(r, c),
	}
}
