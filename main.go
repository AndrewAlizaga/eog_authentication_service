package main

import (
	"fmt"
	"log"

	application "github.com/AndrewAlizaga/eog_authentication_service/app"
)

func main() {
	fmt.Println("Hello, World!")

	response, err := application.StartService(true)

	if err != nil {
		log.Fatal("FAILURE ERROR: ", err)
	}

	fmt.Println(response)

}
