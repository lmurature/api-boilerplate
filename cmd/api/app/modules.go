package app

import (
	"github.com/lmurature/api-boilerplate/cmd/api/clients"
	"github.com/lmurature/api-boilerplate/cmd/api/controllers"
	"github.com/lmurature/api-boilerplate/cmd/api/domain"
	"github.com/lmurature/api-boilerplate/cmd/api/lib"
	"github.com/lmurature/api-boilerplate/cmd/api/routes"
	"github.com/lmurature/api-boilerplate/cmd/api/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	lib.Module,
	runnerModule,
	domain.DaoModule,
	clients.Module,
)
