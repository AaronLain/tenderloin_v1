package main

import (
	//"encoding/csv"
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


type Record struct {
	OrderNum			string `csv:"Name"`
	OrderDate			string `csv:"Created at"`
	DatePaid			string `csv:"Paid at"`
	Total				string `csv:"Total"`
	AmountPaid			string `csv:"Total"`
	Tax					string `csv:"Taxes"`
	ShippingPaid		string `csv:"Shipping"`
	ShippingService		string `csv:"Shipping Method"`
	CustomField1		string 
	CustomField2		string `csv:"Tags"`
	CustomFieldF3		string
	Source				string `csv:"Source"`
	BuyerFullName		string `csv:"Billing Name"`	
	BuyerEmail			string `csv:"Email"`
	BuyerPhone			string `csv:"Billing Phone"`
	RecFullName			string `csv:"Shipping Name"`
	RecPhone			string `csv:"Shipping Phone"`
	RecCompany			string `csv:"Shipping Company"`
	AddressLine1		string `csv:"Shipping Address1"`
	AddressLine2		string `csv:"Shipping Address2"`
	City				string `csv:"Shipping City"`
	State				string `csv:"Shipping Province"`
	PostalCode			string `csv:"Shipping Zip"`
	CountryCode			string `csv:"Shipping Country"`
	ItemSKU				string `csv:"Lineitem sku"`
	ItemName			string `csv:"Lineitem name"`
	ItemUnitPrice		string `csv:"Lineitem price"`
}

// The idea is to keep the keys for each order with their respective zip/temps
type ZipTemp struct {
	Keys []int
	Zip string
	Temp string
}
// Might not be necessary, we will see.
type ZipTempTable struct {
	ZipTemps []ZipTemp
}
		
func main() {	
	csvReader()
}

func firstFiveZip(s string) string {
	i := 0
	for j := range s {
		if i == 5 {
			return s[:j]
		}
		i++
	}
	if len(s) < 5 {
		z := "0"
		s := z + s

		return s
	}

	return s
}

func convertAllZips(r []*Record) []*Record {
	for _, v := range r {		
		zipFiveDig := firstFiveZip(v.PostalCode)
		v.PostalCode = zipFiveDig
	}
	return r
}

func csvReader() {
	// step 1: open file
	recordFile, err := os.Open("./orders_test.csv")
	if err != nil {
		fmt.Println("Error occured! ::", err)
	}
	defer recordFile.Close()

	records := []*Record{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		panic(err)
	}	

	newRecords := convertAllZips(records)
	
	for _, v := range newRecords {
		fmt.Printf("new: %v \n", v.PostalCode)
	}
}
 
