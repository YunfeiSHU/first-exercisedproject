package middleware

import (
	"gin-jwt-gorm/domain/model"
	"gin-jwt-gorm/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuthMidddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		t := strings.Split(tokenString, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, model.ErrorResponse{
						Message: err.Error(),
					})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: err.Error(),
			})
			c.Abort()
			return
		}
	}
}
