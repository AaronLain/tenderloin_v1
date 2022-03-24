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


func csvWriter(input string, o []*orders.OrderRecord) {
	output1 := strings.TrimSuffix(input, ".csv")
	output2 := strings.TrimPrefix(output1, "./")
	outputName := output2 + "_"
	newRecords := zip.GetTemps(o)
	// check to see if filename already exists before creating!
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		file, err := ioutil.TempFile("./", outputName)
		fmt.Printf("file: %v", file.Name())
		if err != nil {
			fmt.Println("Couldn't create csv ::", err)
		}
		gocsv.MarshalFile(&newRecords, file)
	}
}

func initializeCSV() {
	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	records := csvReader(fileName)

	defer zip.GetTemps(records)

}

func main() {
	initialize()
}
