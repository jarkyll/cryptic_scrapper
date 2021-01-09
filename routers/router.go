package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	router.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router = initializeRouterA(router)
	router = initializeRouterB(router)
	router = InitializeAuthenticationRoutes(router)

	return router

}