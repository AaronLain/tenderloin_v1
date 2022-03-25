package main

import (
	orders "ajl/tenderloin/orders"
	zip "ajl/tenderloin/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

func csvReader(s string) ([]*orders.OrderRecord, error) {
	recordFile, err := os.Open(s)
	if err != nil {
		fmt.Println("Reader Error occured! ::", err)
	}

	records := []*orders.OrderRecord{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		fmt.Println("Unmarshalling Error occured! ::", err)
	}
	defer recordFile.Close()

	return records, err
}

func csvWriter(input string, o []*orders.OrderRecord) {
	output1 := strings.TrimSuffix(input, ".csv")
	output2 := strings.TrimPrefix(output1, "./")
	outputName := output2 + "_"
	newRecords, err := zip.GetTemps(o)
	if err != nil {
		fmt.Println("Failed to get new records ::", err)
	}
	// check to see if filename already exists before creating
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		//this puts the csv in the local file
		file, err := ioutil.TempFile("./", outputName)
		if err != nil {
			fmt.Println("Can't create csv ::", err)
		}
		gocsv.MarshalFile(&newRecords, file)
	}
}

func initializeCSV() {
	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	records, err := csvReader(fileName)
	if err != nil {
		fmt.Println("Can't initialize reader ::", err)
	}

	csvWriter(fileName, records)

}

func main() {
	initializeCSV()
}
