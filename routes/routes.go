package routes

import (
	"cashapp/infra"
	"cashapp/services"

	"github.com/gin-gonic/gin"
)

type router struct {
	engine   *gin.Engine
	config   *infra.Config
	services services.Services
}

func NewRouter(engine *gin.Engine, config *infra.Config, services services.Services) *router {
	return &router{
		engine:   engine,
		config:   config,
		services: services,
	}
}

func (r *router) RegisterRoutes() {

}
