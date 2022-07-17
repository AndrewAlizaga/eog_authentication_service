package routes

import (
	"fmt"
	"os"

	handlers "github.com/AndrewAlizaga/eog_authentication_service/app/handlers"
	"github.com/gin-gonic/gin"
)

func MainRouter() {
	router := gin.Default()

	port := os.Getenv("PORT")

	fmt.Println("using port: ", port)

	//ROUTE GROUPS
	AuthRouter(router)
	SearchRouter(router)

	//ACCOUNT GROUPS
	account := router.Group("/accounts")
	{
		//Get Accounts
		account.GET("/accounts", handlers.GetAccounts)

		//Get Account
		account.GET("/accounts/:id", handlers.GetAccountById)

		//Post Account
		account.POST("/accounts", handlers.PostAccount)

		//Update Account
		account.PUT("/accounts/:id", handlers.Auth, handlers.UpdateAccount)

	}

	router.Run(":" + port)
}
