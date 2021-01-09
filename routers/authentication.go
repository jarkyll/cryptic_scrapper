package routers

import (
	"github.com/gin-gonic/gin"
	"golang-boilerplate/controllers"
)

func InitializeAuthenticationRoutes(router *gin.Engine) *gin.Engine {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.HandleRegister)
		authRouter.POST("/login", controllers.HandleLogin)
		authRouter.POST("/logout", controllers.HandleLogout)
		authRouter.POST("/refresh-token", controllers.HandleRefreshToken)
	}

	return router;
}
