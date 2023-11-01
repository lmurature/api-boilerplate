package config

import (
	"github.com/gin-contrib/cors"
	"math"
	"os"
	"time"
)

const (
	AccessTokenExpirationMs    = 3600000
	RefreshTokenExpirationNano = math.MaxInt64
)

var (
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
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

func GetDbHost() string {
	return dbHost
}

func GetDbName() string {
	return dbName
}

func GetDbUser() string {
	return dbUser
}

func GetDbPass() string {
	return dbPass
}
