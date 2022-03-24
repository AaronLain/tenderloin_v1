package main

import (
	orders "ajl/tenderloin/orders"
	zip "ajl/tenderloin/zip"
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

func csvReader() {

	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	recordFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occured! ::", err)
	}
	defer recordFile.Close()

	records := []*orders.OrderRecord{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		panic(err)
	}

	zip.GetTemps(records)

}

func main() {
	csvReader()
}
