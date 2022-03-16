package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvReader()
}	

func csvReader() {
	// step 1: open file
	recordFile, err := os.Open("./orders_test.csv")
	if err != nil {
		fmt.Println("Error occured! ::", err)
	}

	// step 2: initialize reader
	reader := csv.NewReader(recordFile)

	// step 3: read all records
	records, _ := reader.ReadAll()


	fmt.Println(records)
}
 
