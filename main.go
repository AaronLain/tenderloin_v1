package main

import (
	//"encoding/csv"
	zip "ajl/tenderloin/zip"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	//	"unicode/utf8"
)

// Shipsitation requires the following fields:
// Order Number, Order Date, Date Paid, Order Total,
// Amount Paid, Tax, Shipping Paid, Shipping Service, Weight (oz),
// Box Height, Box Width, Box Length, Custom Field 1, Custom Field 2, Custom Field 3, Order Source
// Additional: Email,

func main() {
	csvReader()
}

func csvReader() {

	recordFile, err := os.Open("./orders_test.csv")
	if err != nil {
		fmt.Println("Error occured! ::", err)
	}
	defer recordFile.Close()

	records := []*zip.Record{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		panic(err)
	}

	newRecords := zip.ConvertAllZips(records)

	for _, v := range newRecords {
		fmt.Printf("new: %v \n", v.PostalCode)
	}

	zip.PrintHello()
}
