package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	Claim "github.com/AndrewAlizaga/eog_authentication_service/app/models/claim"

	db "github.com/AndrewAlizaga/eog_authentication_service/app/db"
	Account "github.com/AndrewAlizaga/eog_authentication_service/app/models/account"
	jwt "github.com/AndrewAlizaga/eog_authentication_service/app/security/access"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func Login(c *gin.Context) {

	println("reached login function")
	//account, err
	user, password, hasAuth := c.Request.BasicAuth()

	fmt.Println("user: ", user)
	fmt.Println("password: ", password)
	fmt.Println("hasAuth: ", hasAuth)

	//First catch block
	if hasAuth && user != "" && password != "" {

		//find user on db
		result, err := db.FindOne("account", bson.M{"email": user})

		if err != nil {

			print("user not found or server error")
			errorResponse := struct {
				Message string `json:"message"`
			}{"server error"}

			c.IndentedJSON(http.StatusBadRequest, errorResponse)
			return
		}

		fmt.Println("user found, proceed password validation")
		var user Account.Account

		result.Decode(&user)

		if user == (Account.Account{}) {
			print("user not found or server error")
			errorResponse := struct {
				Message string `json:"message"`
			}{"user not found"}

			c.IndentedJSON(http.StatusBadRequest, errorResponse)
			return
		}

		fmt.Println("[DEBUG] USER BODY: ", user)

		fmt.Println("[DEBUG] USER ID: ", user.Id)

		match := user.ComparePassword(password)

		if !match {

			print("user credentials dont match")
			errorResponse := struct {
				Message string `json:"message"`
			}{"credentials error"}

			c.IndentedJSON(http.StatusBadRequest, errorResponse)
			return
		}

		userStruct := Claim.ClaimUserObj{
			Name:  user.Name,
			Email: user.Email,
			Id:    user.Id,
		}

		fmt.Println("user: ", user.Name)
		fmt.Println("user strcuted: ", userStruct)
		//Proceed with user finding,
		//and jwt creation
		jwt, err := jwt.NewToken(userStruct, time.Now())

		fmt.Println("Current obtained token: ", jwt)

		if err != nil {
			c.String(http.StatusForbidden, "Error creating token, reason: ", err)
			c.Abort()
			return
		}

		response := struct {
			Jwt string `json:"jwt"`
		}{jwt}

		fmt.Println("jwt response: ", response)

		c.IndentedJSON(http.StatusOK, response)

	} else {
		c.String(http.StatusForbidden, "No authorization header provided")
		c.Abort()
		return
	}

}

func Auth(c *gin.Context) {

	const BEARER = "Bearer"
	fmt.Println("arrived the auth function")

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		fmt.Printf("NIL HEADER")
		c.IndentedJSON(http.StatusBadRequest, errors.New("no authorization header").Error())
		return
	}
	fmt.Println(authHeader)

	fmt.Println("getting bearer")
	tokenData := authHeader[len(BEARER)+1:]

	fmt.Println(tokenData)

	result, payload, err := jwt.ValidateToken(tokenData)

	fmt.Println("[DEBUG] token validated")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	/*
		userStruct := struct {
			name  string
			email string
			id    string
		}{
			user.Name,
			user.Email,
			user.Id,
		} */

	if result {

		//if result attempt to decode payload into req object

		claimObj, ok := payload.(*Claim.Claim)

		if !ok {
			fmt.Println("NOT OK")
			fmt.Println(payload)
			fmt.Println(claimObj)
			c.IndentedJSON(http.StatusBadRequest, errors.New("invalid token"))
			return
		}

		fmt.Println("[CONVERTED]: CLAIM")
		fmt.Println(claimObj)

		fmt.Println("[DEBUG] AFTER ATTEMPTING TO SET CLAIM")
		c.Set("userClaim", claimObj.Data)

		fmt.Println("USER CLAIM SET")
		fmt.Println("CLAIM IS: ", payload)

		c.IndentedJSON(http.StatusAccepted, result)
		c.Next()
	} else {
		c.IndentedJSON(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

}
