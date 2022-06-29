package routes

import (
	handlers "github.com/AndrewAlizaga/eog_authentication_service/app/handlers"
	"github.com/gin-gonic/gin"
)

//auth router definition
func AuthRouter(router *gin.Engine) {

	AuthRouter := router.Group("/auth")
	{
		AuthRouter.POST("/login", handlers.Login)
		AuthRouter.GET("/", handlers.Auth)
	}
}
