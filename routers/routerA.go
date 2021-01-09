package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRouterA(router *gin.Engine) *gin.Engine {
	example := router.Group("/test")
	{
		example.GET("pong", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		example.GET("ponger", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}

	return router
}