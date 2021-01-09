package middlewares

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang-boilerplate/service"
	"net/http"
)

func AuthsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.Request

		if service.IsTokenValid(request) {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}