package auth_utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lmurature/api-boilerplate/cmd/api/config"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"time"
)

const (
	TokenTypeRefresh = "REFRESH_TOKEN"
	TokenTypeAccess  = "ACCESS_TOKEN"

	Bearer = "bearer"
)

func GenerateToken(userId int64, tokenType string) (*string, apierrors.ApiError) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	var expirationTime int64
	if tokenType == TokenTypeRefresh {
		expirationTime = time.Now().Add(time.Duration(config.RefreshTokenExpirationNano)).Unix()
	} else if tokenType == TokenTypeAccess {
		expirationTime = time.Now().Add(time.Millisecond * time.Duration(config.AccessTokenExpirationMs)).Unix()
	}
	claims["exp"] = expirationTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(err.Error(), err)
	}

	return &signed, nil
}
