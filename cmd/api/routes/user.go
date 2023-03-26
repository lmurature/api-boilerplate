package routes

import (
	"github.com/lmurature/api-boilerplate/cmd/api/controllers"
	"github.com/lmurature/api-boilerplate/cmd/api/lib"
	"github.com/lmurature/api-boilerplate/cmd/api/middlewares"
)

type UserRoutes struct {
	handler         *lib.RequestHandler
	usersController *controllers.UsersController
}

func NewUserRoutes(handler *lib.RequestHandler, controller *controllers.UsersController) *UserRoutes {
	return &UserRoutes{
		handler:         handler,
		usersController: controller,
	}
}

func (u UserRoutes) Setup() {
	u.handler.Gin.POST("/users/create", u.usersController.RegisterUser)
	u.handler.Gin.POST("/users/auth/generate_token", u.usersController.AuthenticateUser)
	u.handler.Gin.POST("/users/auth/refresh_token", u.usersController.RefreshUserToken)
	u.handler.Gin.GET("/users/:user_id", middlewares.Authenticate, u.usersController.GetUser)
}
