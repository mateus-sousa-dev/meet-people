package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-sousa-dev/meet-people/internal/auth"
	"net/http"
)

func Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		next(c)
	}
}
