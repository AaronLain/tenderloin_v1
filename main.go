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

func csvReader(s string) []*orders.OrderRecord {
	recordFile, err := os.Open(s)
	if err != nil {
		fmt.Println("Couldn't open ::", err)
	}

	records := []*orders.OrderRecord{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		panic(err)
	}
	defer recordFile.Close()

	return records
}

func csvWriter(input string, o []*orders.OrderRecord) {
	count := 1
	num := fmt.Sprintf("%v", count)
	output1 := strings.TrimSuffix(input, ".csv")
	output2 := strings.TrimPrefix(output1, "./")
	outputName := output2 + "_" + num
	// fmt.Printf("OutputName %v", outputName)
	newRecords := zip.GetTemps(o)
	// fmt.Printf(" write new records?: %v \n", newRecords)
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		file, err := ioutil.TempFile("./", outputName)
		fmt.Printf("file: %v", file.Name())
		if err != nil {
			fmt.Println("Couldn't create csv ::", err)
		}
		gocsv.MarshalFile(&newRecords, file)
	}

}

func initialize() {
	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	records := csvReader(fileName)

	defer csvWriter(fileName, records)

}

func main() {
	initialize()
}
