package handlers

import (
	"errors"
	"fmt"
	"net/http"

	Account "github.com/AndrewAlizaga/eog_authentication_service/app/models/account"
	"github.com/AndrewAlizaga/eog_authentication_service/app/security/encryption"
	serviceModels "github.com/AndrewAlizaga/eog_authentication_service/app/services/models"
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

	var accountUpdate Account.AccountRequest

	accountParameter := c.Param("id")

	if accountParameter == "" {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	fmt.Println("PARAMETER: ", accountParameter)
	err := c.BindJSON(&accountUpdate)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	fmt.Println("ACCOUNT REQUEST: ", accountUpdate)

	result, err := serviceModels.UpdateAccount(accountParameter, accountUpdate)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusAccepted, result)
}
