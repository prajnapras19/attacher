package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prajnapras19/attacher/constants"
	"github.com/prajnapras19/attacher/lib"
	"github.com/prajnapras19/attacher/user"
)

var (
	ErrUnauthorizedRequest = errors.New("unauthorized request")
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func JWTTokenMiddleware(userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(constants.Token)
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, lib.BaseResponse{
				Message: ErrUnauthorizedRequest.Error(),
			})
			c.Abort()
			return
		}
		claims, err := userService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, lib.BaseResponse{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set(constants.JWTClaims, claims)

		c.Next()
	}
}
