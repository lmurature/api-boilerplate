package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lmurature/api-boilerplate/cmd/api/config"
	"github.com/lmurature/api-boilerplate/cmd/api/domain/apierrors"
	"net/http"
	"strconv"
	"strings"
)

func Authenticate(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")

	if reqToken == "" {
		apierror := apierrors.NewForbiddenApiError("Authorization token not provided")
		c.JSON(http.StatusForbidden, apierror)
		c.Abort()
		return
	}

	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		err := apierrors.NewBadRequestApiError("authorization token (Bearer) is needed to access this endpoint")
		c.JSON(err.Status(), err)
		return
	}

	token := splitToken[1]

	jwtToken, jwtErr := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetAuthKey()), nil
	})
	if jwtErr != nil {
		apiErr := apierrors.NewUnauthorizedApiError(jwtErr.Error())
		c.JSON(apiErr.Status(), apiErr)
		c.Abort()
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		userId, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if err != nil {
			apiErr := apierrors.NewInternalServerApiError("invalid users id to parse token", err)
			c.JSON(apiErr.Status(), apiErr)
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
