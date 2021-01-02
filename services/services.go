package services

import (
	"cashapp/core"
	"cashapp/repo"

	"github.com/go-redis/redis/v8"
)

type Services struct {
	Accounts *accountLayer
}

func NewService(r repo.Repo, kvStore *redis.Client, c *core.Config) Services {
	return Services{
		Accounts: newAccountLayer(r, c),
	}
}
