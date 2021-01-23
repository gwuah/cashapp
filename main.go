package main

import (
	"cashapp/core"
	"cashapp/repository"
	"cashapp/routes"
	"cashapp/services"

	"cashapp/core/db"
	"cashapp/models"

	"log"
)

func main() {
	config := core.NewConfig()

	pg, err := db.NewPostgres(config)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
	}

	err = db.RunMigrations(pg, &models.Transaction{}, &models.Account{}, &models.Wallet{})
	if err != nil {
		log.Fatal("failed to run migrations. err:", err)
	}

	if config.RUN_SEEDS {
		models.RunSeeds(pg)
	}

	cache := db.NewRedis(config)
	repo := repository.NewRepository(pg)
	service := services.NewService(repo, cache, config)

	server := core.NewHTTPServer(config)
	router := routes.NewRouter(server.Engine, config, service)

	router.RegisterRoutes()
	server.Start()

}
