package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/prajnapras19/attacher/constants"
)

// TODO: add more claims
type JWTClaims struct {
	jwt.StandardClaims
	ID       uint   `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Serial   string `json:"serial" binding:"required"`
}

func GetJWTClaimsFromContext(c *gin.Context) (*JWTClaims, error) {
	if val, exists := c.Get(constants.JWTClaims); exists {
		if res, ok := val.(*JWTClaims); ok {
			return res, nil
		}
		return nil, ErrFailedToParseJWTClaimsInContext
	}
	return nil, ErrJWTClaimsNotFoundInContext
}
