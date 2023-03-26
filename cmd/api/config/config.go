package config

import (
	"github.com/gin-contrib/cors"
	"math"
	"time"
)

const (
	AccessTokenExpirationMs    = 3600000
	RefreshTokenExpirationNano = math.MaxInt64
)

func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}
}
