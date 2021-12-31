package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/anvari1313/splitwise.go"
	"github.com/joho/godotenv"
)

func ParseCost(s string) string {
	v, err := strconv.ParseFloat(s, 32)

	if err != nil {
		log.Fatalln(err)
	}

	if v < 0 {
		v = v * -1
		return fmt.Sprintf("%f", v)
	} else {
		fmt.Println("Log: cost is", s)
	}

	return s
}

func ParseDate(s string) string {
	/*
		ParseDate("12/19/2021")
		-> "2021-12-19"
	*/
	t, err := time.Parse("01/02/2006", s)

	if err != nil {
		log.Println("WARNING: ", err)
	} else {
		return t.Format("2006-01-02")
	}

	t, err = time.Parse("2006-01-02", s)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	return t.Format("2006-01-02")

}

func csvReader() []splitwise.ExpenseDTO {
	/*
		line format Date,Description,Cost,
			12/20/2021,Audible*173OK5SQ3,7.95,
	*/
	// reads file from command line
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// convert file to csv
	csvFile, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// creates slice of ExpenseDTO
	csvLines := []splitwise.ExpenseDTO{}

	// parse argument to uint64
	group_id, _ := strconv.ParseUint(os.Args[2], 10, 64)

	// loop through csv file
	for _, line := range csvFile {
		// fmt.Println(ParseDate(line[0]))
		// convert to struct
		csvLine := splitwise.ExpenseDTO{
			GroupId:      group_id,
			Date:         ParseDate(line[0]),
			SplitEqually: true,
			Description:  line[1],
			Cost:         ParseCost(line[2]),
			Details:      "Added by LeoBot™",
		}

		fmt.Println(csvLine)

		// append to slice
		csvLines = append(csvLines, csvLine)

	}
	// return nil
	return csvLines
}

var api_key string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln(err)
	}
	api_key = os.Getenv("SPLITWISE_API_KEY")

	log.Println(api_key)
}

func main() {
	auth := splitwise.NewAPIKeyAuth(api_key)
	client := splitwise.NewClient(auth)

	for _, expense := range csvReader() {
		e, err := client.CreateExpense(context.Background(), &expense)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(e)
	}
}

// _ = csvReader()
// group_id, _ := strconv.ParseUint(os.Args[2], 10, 64)
// e := splitwise.ExpenseDTO{
// 	GroupId:      group_id,
// 	SplitEqually: true,
// 	Cost:         "100",
// 	Description:  "foo",
// 	Date:         "2012-05-02",
// 	Details:      "Added by LeoBot™",
// }
// client.CreateExpense(context.Background(), &e)
// if err != nil {
// 	panic(err)
// }
// ex, err := json.MarshalIndent(expense, "", "    ")
// fmt.Println(string(ex))
// fmt.Println(expense)

// group_id, _ := strconv.ParseUint(os.Args[2], 10, 64)
// group, err := client.GroupByID(context.Background(), group_id)
// if err != nil {
// 	panic(err)
// }

// b, err := json.MarshalIndent(group, "", "    ")
// if err != nil {
// 	fmt.Println("error:", err)
// }
// fmt.Println(string(b))

// GET Expenses
// e, _ := client.Expenses(context.Background())

// parsedExpenses, _ := json.MarshalIndent(e, "", "    ")
// fmt.Println(string(parsedExpenses))

// }
