package lib

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lmurature/api-boilerplate/cmd/api/config"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() *RequestHandler {
	engine := gin.Default()
	engine.Use(cors.New(config.GetCorsConfig()))
	return &RequestHandler{Gin: engine}
}
