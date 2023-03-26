package routes

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewHealthRoutes), fx.Provide(NewUserRoutes), fx.Provide(NewRoutes))

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(userRoutes *UserRoutes, healthRoutes *HealthRoutes) Routes {
	return Routes{
		userRoutes,
		healthRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
