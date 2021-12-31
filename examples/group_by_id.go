package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/anvari1313/splitwise.go"
	"github.com/joho/godotenv"
)

var api_key string
var group_id uint64

func init() {
	godotenv.Load("../.env")
	api_key = os.Getenv("SPLITWISE_API_KEY")
	flag.Uint64Var(&group_id, "group_id", 0, "the group_id")
}

func main() {
	auth := splitwise.NewAPIKeyAuth(api_key)
	if len(api_key) == 0 {
		log.Fatalln("auth_key is empty or .env file not present")
	}

	client := splitwise.NewClient(auth)

	if group_id == 0 {
		log.Println("group_id not passed... defaulting value to 0")
	}

	group, err := client.GroupByID(context.Background(), group_id)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.MarshalIndent(group, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

}
