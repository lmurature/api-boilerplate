package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/users"
	"github.com/lmurature/api-boilerplate/cmd/api/services"
	"net/http"
	"strconv"
)

const (
	currentUser = "me"
)

type UsersController struct {
	usersService services.UsersServiceInterface
}

func NewUserController(usersService *services.UsersServiceInterface) *UsersController {
	return &UsersController{usersService: *usersService}
}

func (u *UsersController) RegisterUser(c *gin.Context) {
	var request users.RegisterUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError("error in register users json body format"))
		return
	}

	response, err := u.usersService.RegisterUser(context.TODO(), request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (u *UsersController) AuthenticateUser(c *gin.Context) {
	var request users.AuthenticateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError("error in authenticate users json body format"))
		return
	}

	response, err := u.usersService.AuthenticateUser(context.TODO(), request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (u *UsersController) RefreshUserToken(c *gin.Context) {
	var request users.RefreshUserTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError("error in authenticate users json body format"))
		return
	}

	response, err := u.usersService.RefreshUserToken(context.TODO(), request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (u *UsersController) GetUser(c *gin.Context) {
	var userId int64
	if c.Param("user_id") == currentUser {
		callerId, _ := c.Get("user_id")
		userId = callerId.(int64)
	} else {
		userId, _ = strconv.ParseInt(c.Param("user_id"), 10, 64)
	}
	resp, err := u.usersService.GetUserByID(context.TODO(), userId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
