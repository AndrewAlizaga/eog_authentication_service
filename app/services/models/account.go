package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/AndrewAlizaga/eog_authentication_service/app/db"
	Account "github.com/AndrewAlizaga/eog_authentication_service/app/models/account"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// dummy account data.
var accounts = []Account.Account{
	{
		Name:           "Joe Duffin",
		AccountStatus:  "BLOCKED",
		RecentActivity: "ACCOUNT_BLOCKED",
		Password:       "qwerty",
		CreationDate:   time.Now(),
		Email:          "joeduffin1@gmail.com",
	},
	{
		Name:           "Joe Duffin",
		AccountStatus:  "BLOCKED",
		RecentActivity: "ACCOUNT_BLOCKED",
		Password:       "qwerty",
		CreationDate:   time.Now(),
		Email:          "joeduffin13@gmail.com",
	},
	{
		Name:           "Joe Duffin",
		AccountStatus:  "BLOCKED",
		RecentActivity: "ACCOUNT_BLOCKED",
		Password:       "qwerty",
		CreationDate:   time.Now(),
		Email:          "joeduffin12@gmail.com",
	},
}

func CheckPassword(account Account.Account, password string) (bool, error) {

	//encrypt password

	//check passwords

	//solving return
	return true, nil
}

func GetAccounts() ([]Account.Account, error) {

	return accounts, nil
}

func GetAccount(accountId string) (Account.Account, error) {

	var account Account.Account

	result, err := db.FindOne("account", bson.M{"_id": accountId})

	fmt.Println("ACCOUNT OBTAINED")
	if err != nil {
		return account, err
	}

	fmt.Println(result)

	result.Decode(&account) // (Account.Account)

	fmt.Println("account converted")
	if account == (Account.Account{}) {
		return account, errors.New("account not found")
	}

	return account, nil

}

func CreateAccount(account Account.Account) (string, error) {

	//Check email is not on use
	existingAccountResult, err := db.FindOne("account", bson.M{"email": account.Email})

	if err != nil {
		return "", err
	}

	var previousAccount Account.Account

	fmt.Println("accounts results : ", existingAccountResult)
	existingAccountResult.Decode(&previousAccount)
	fmt.Println("decoded  results: ", previousAccount)
	if previousAccount != (Account.Account{}) {
		fmt.Println("already on use")
		return "", errors.New("email already on use")
	}

	result, err := db.Insert("account", account)

	if err != nil {
		return "", err
	}

	fmt.Println(result)

	InsertionResult := result.InsertedID.(primitive.ObjectID).Hex()

	fmt.Println("conveted interface nested obj to string")

	return InsertionResult, nil

}

func UpdateAccount(accountId string, update Account.AccountRequest) (string, error) {

	objID, err := primitive.ObjectIDFromHex(accountId)

	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println(objID)
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	updateQuery := bson.M{"$set": update}

	//updateQuery := bson.M{"$set": string(response)}

	fmt.Println(updateQuery)

	result, err := db.Update("account", filter, updateQuery)

	fmt.Println("POST DB QUERY")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if result.ModifiedCount > 0 {
		return fmt.Sprint(result.ModifiedCount), nil
	}

	return "", errors.New("documents not updated")
}
