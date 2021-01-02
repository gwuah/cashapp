package infra

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development Environment = "dev"
	Staging     Environment = "staging"
)

type Config struct {
	PG_HOST        string
	PG_PORT        string
	PG_NAME        string
	PG_USER        string
	PG_PASS        string
	PG_SSLMODE     string
	REDIS_ADDRESS  string
	REDIS_PASSWORD string
	REDIS_DB       int
	REDIS_URL      string
	DATABASE_URL   string
	PORT           int
	RUN_SEEDS      bool
	ENVIRONMENT    Environment
}

func Get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("%s: %s", key, err)
			return fallback
		}
		return i
	}
	return fallback
}

func GetEnvironment() Environment {
	if env := Get("ENV", ""); env == "" {
		return Development
	} else {
		return Environment(env)
	}
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading env file")
	}
	return &Config{
		PG_HOST:        Get("PG_HOST", "localhost"),
		PG_PORT:        Get("PG_PORT", "5432"),
		PG_NAME:        Get("PG_NAME", "cashapp"),
		PG_USER:        Get("PG_USER", "user"),
		PG_PASS:        Get("PG_PASS", "password"),
		PG_SSLMODE:     Get("PG_SSLMODE", "disable"),
		REDIS_ADDRESS:  Get("REDIS_ADDRESS", "localhost:6379"),
		REDIS_PASSWORD: Get("REDIS_PASSWORD", ""),
		REDIS_DB:       GetInt("REDIS_DB", 1),
		REDIS_URL:      Get("REDIS_URL", ""),
		DATABASE_URL:   Get("DATABASE_URL", ""),
		PORT:           GetInt("PORT", 5454),
		ENVIRONMENT:    GetEnvironment(),
		RUN_SEEDS:      true,
	}
}
