package db

import (
	"cashapp/infra"
	"net/url"

	"github.com/go-redis/redis/v8"
)

func NewRedis(config *infra.Config) *redis.Client {
	if infra.GetEnvironment() == infra.Staging {
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
