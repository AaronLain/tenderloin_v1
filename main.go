package main

import (
	zip "ajl/tenderloin/zip"
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

	records := []*zip.Record{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		panic(err)
	}

	zip.CreateRawZipTable(records)
}

func main() {
	csvReader()
}
