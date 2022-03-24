package main

import (
	orders "ajl/tenderloin/orders"
	zip "ajl/tenderloin/zip"
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

func csvReader(s string) []*orders.OrderRecord {
	recordFile, err := os.Open(s)
	if err != nil {
		fmt.Println("Error occured! ::", err)
	}

	records := []*orders.OrderRecord{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		panic(err)
	}
	defer recordFile.Close()

	return records
}

func initialize() {
	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	records := csvReader(fileName)

	defer zip.GetTemps(records)

}

func main() {
	initialize()
}
