package app

import (
	"errors"
	"fmt"

	"github.com/AndrewAlizaga/eog_authentication_service/app/db"
	routes "github.com/AndrewAlizaga/eog_authentication_service/app/routes"
)

func StartService(start bool) (string, error) {

	//KICK OFF MONGODB
	_, err := db.GetMongoClient()

	if err != nil {
		fmt.Println("error connecting")
		panic(err)
	}

	//KICK OFF GIN SERVER
	routes.MainRouter()

	if !start {
		return "", errors.New("start is off")
	}

	fmt.Printf("STARTING AUTH SERVICE")
	message := fmt.Sprintf("Starting, %v.", start)

	return message, nil
}
