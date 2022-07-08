package routes

import (
	handlers "github.com/AndrewAlizaga/eog_authentication_service/app/handlers"
	"github.com/gin-gonic/gin"
)

//search router definition
func SearchRouter(router *gin.Engine) {

	SearchRouter := router.Group("/search")
	{
		SearchRouter.POST("/search", handlers.Login)
		SearchRouter.GET("/", handlers.Auth)
	}
}
