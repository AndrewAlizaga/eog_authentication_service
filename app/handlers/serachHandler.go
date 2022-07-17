package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	Account "github.com/AndrewAlizaga/eog_authentication_service/app/models/account"
	"github.com/AndrewAlizaga/eog_authentication_service/app/security/encryption"
	serviceModels "github.com/AndrewAlizaga/eog_authentication_service/app/services/models"
	searchv1 "github.com/AndrewAlizaga/eog_protos/pkg/proto/search"
	searchService "github.com/AndrewAlizaga/grpc_client_eog_go/pkg/search"
	"github.com/gin-gonic/gin"
)

func GetSearches(c *gin.Context) {
	//get accounts

	accounts, err := serviceModels.GetAccounts()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, accounts)
}

func GetSearchById(c *gin.Context) {

	id := c.Param("id")

	account, err := serviceModels.GetAccount(id)
	//serviceModels.GetAccount(id)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, account)

}

func GetCaseById(c *gin.Context) {

	var accountRequest Account.AccountRequest

	//Check for bad form
	if err := c.BindJSON(&accountRequest); err != nil {
		fmt.Println("bad request error on body")
		c.IndentedJSON(http.StatusBadRequest, errors.New("body was malformed").Error())
		return
	}

	newAccount := Account.Account{
		Name:     accountRequest.Name,
		Password: encryption.Encrypt(accountRequest.Password),
		Email:    accountRequest.Email,
	}

	result, err := serviceModels.CreateAccount(newAccount)

	fmt.Println("Post db insertion")

	if err != nil {
		fmt.Println(err)
		fmt.Println("on use error return")
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func PostSearch(c *gin.Context) {

	var searchPost searchv1.Search

	if err := c.BindJSON(&searchPost); err != nil {
		log.Print("ERROR BINDING BODY JSON")
		log.Print(err.Error())
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}


	fmt.Println("SEARCH REQUEST: ", searchPost)

	result, err := searchService.PostSearch(&searchPost)

	//result, err := serviceModels.UpdateAccount(accountParameter, searchPost)
	log.Println(result)
	if err != nil {
		log.Println("ERROR CALLING EOG: ", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusAccepted, result)
}
