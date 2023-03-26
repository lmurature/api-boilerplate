package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
