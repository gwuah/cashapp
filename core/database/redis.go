package database

import (
	"cashapp/core"
	"net/url"

	"github.com/go-redis/redis/v8"
)

func NewRedis(config *core.Config) *redis.Client {
	if core.GetEnvironment() == core.Staging {
		parsedURL, _ := url.Parse(config.REDIS_URL)
		password, _ := parsedURL.User.Password()
		return redis.NewClient(&redis.Options{
			Addr:     parsedURL.Host,
			Password: password,
		})
	}

	return redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDRESS,
		Password: config.REDIS_PASSWORD,
		DB:       config.REDIS_DB,
	})
}
