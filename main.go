package main

import (
	"cashapp/core"
	"cashapp/core/database"

	"cashapp/repository"
	"cashapp/routes"
	"cashapp/services"

	"cashapp/models"

	"log"
)

func main() {
	config := core.NewConfig()

	pg, err := database.NewPostgres(config)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
	}

	err = database.RunMigrations(pg, &models.Transaction{}, &models.User{}, &models.Wallet{})
	if err != nil {
		log.Fatal("failed to run migrations. err:", err)
	}

	if config.RUN_SEEDS {
		models.RunSeeds(pg)
	}

	cache := database.NewRedis(config)
	repo := repository.NewRepository(pg)
	service := services.NewService(repo, cache, config)

	server := core.NewHTTPServer(config)
	router := routes.NewRouter(server.Engine, config, service)

	router.RegisterRoutes()
	server.Start()

}
