package app

import (
	"github.com/lmurature/api-boilerplate/cmd/api/lib"
	"github.com/lmurature/api-boilerplate/cmd/api/routes"
	"go.uber.org/fx"
)

type Runner struct{}

func NewRunner(h *lib.RequestHandler, r routes.Routes) *Runner {
	r.Setup()
	h.Gin.Run(":8080")
	return &Runner{}
}

var runnerModule = fx.Provide(NewRunner)
