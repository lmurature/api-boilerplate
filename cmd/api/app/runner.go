package app

import (
	"fmt"
	"github.com/lmurature/api-boilerplate/cmd/api/config"
	"github.com/lmurature/api-boilerplate/cmd/api/lib"
	"github.com/lmurature/api-boilerplate/cmd/api/routes"
	"go.uber.org/fx"
)

type Runner struct{}

func NewRunner(h *lib.RequestHandler, r routes.Routes) *Runner {
	r.Setup()
	h.Gin.Run(fmt.Sprintf(":%s", config.GetPort()))
	return &Runner{}
}

var runnerModule = fx.Provide(NewRunner)
