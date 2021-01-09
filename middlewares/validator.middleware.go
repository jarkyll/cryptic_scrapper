package middlewares

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "golang-boilerplate/service"
	"net/http"
)

func ValidatorMiddleware(schema interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := c.ShouldBindJSON(&schema); err != nil {
			// TODO print err somewhere
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			c.Set("request", schema)
			c.Next()
		}
	}
}