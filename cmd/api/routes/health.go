package routes

import (
	"github.com/lmurature/api-boilerplate/cmd/api/controllers"
	"github.com/lmurature/api-boilerplate/cmd/api/lib"
)

type HealthRoutes struct {
	handler          *lib.RequestHandler
	healthController *controllers.HealthController
}

func NewHealthRoutes(handler *lib.RequestHandler, controller *controllers.HealthController) *HealthRoutes {
	return &HealthRoutes{
		handler:          handler,
		healthController: controller,
	}
}

func (h HealthRoutes) Setup() {
	h.handler.Gin.GET("/ping", h.healthController.Ping)
}
