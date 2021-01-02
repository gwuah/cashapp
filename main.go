package main

import (
	"cashapp/infra"
	"cashapp/repo"
	"cashapp/routes"
	"cashapp/services"

	"cashapp/infra/db"
	"cashapp/infra/models"

	"log"
)

func main() {
	config := infra.NewConfig()

	pg, err := db.NewPostgres(config)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
	}

	err = db.RunMigrations(pg, &models.Transaction{}, &models.Account{})
	if err != nil {
		log.Fatal("failed to run migrations. err:", err)
	}

	if config.RUN_SEEDS {
		models.RunSeeds(pg)
	}

	cache := db.NewRedis(config)
	repository := repo.NewRepo(pg)
	server := infra.NewHTTPServer(config)
	service := services.NewService(repository, cache, config)
	router := routes.NewRouter(server.Engine, config, service)
	router.RegisterRoutes()
	server.Start()

}
