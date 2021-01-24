package routes

import (
	"cashapp/core"
	"cashapp/services"

	"github.com/gin-gonic/gin"
)

type router struct {
	engine   *gin.Engine
	config   *core.Config
	services services.Services
}

func NewRouter(engine *gin.Engine, config *core.Config, services services.Services) *router {
	return &router{
		engine:   engine,
		config:   config,
		services: services,
	}
}

func (r *router) RegisterRoutes() {
	RegisterUserRoutes(r.engine, r.services)
	RegisterPaymentRoutes(r.engine, r.services)
}
