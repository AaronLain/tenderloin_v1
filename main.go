package main

import (
	orders "ajl/tenderloin/orders"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func csvReader() {

	recordFile, err := os.Open("./Orders_test_2.csv")
	if err != nil {
		fmt.Println("Error occured! ::", err)
	}
	defer recordFile.Close()

	records := []*orders.OrderRecord{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		panic(err)
	}

	// cleanRecords := zip.CreateZipTable(records)

	orders.AddLineItems(records)
}

func main() {
	csvReader()
}
