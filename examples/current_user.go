package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/anvari1313/splitwise.go"
	"github.com/joho/godotenv"
)

var api_key string

func init() {

	godotenv.Load(".env")

	api_key = os.Getenv("SPLITWISE_API_KEY")
	if len(api_key) == 0 {
		log.Fatalln("auth_key is empty or .env file not present")
	} else {
		fmt.Println("api_key: ", api_key)
	}

}

func main() {
	auth := splitwise.NewAPIKeyAuth(api_key)
	client := splitwise.NewClient(auth)
	currentUser, err := client.CurrentUser(context.Background())
	if err != nil {
		panic(err)
	}
	user, _ := json.MarshalIndent(currentUser, "", "    ")

	fmt.Println(string(user))
}
