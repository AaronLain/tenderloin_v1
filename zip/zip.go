package zip

import (
	o "ajl/tenderloin/orders"
	"fmt"
)

// The idea is to keep the keys for each order with their respective zip/temps
type ZipTemp struct {
	Keys []int
	Zip  string
	Temp string
}

// Might not be necessary, we will see.
type ZipTempTable struct {
	ZipTemps []ZipTemp
}

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}

// keeps only base zip
func FirstFiveZip(s string) string {
	counter := 0
	for i := range s {
		if i == 5 {
			return s[:i]
		}
		counter++
	}
	// Adds a zero to NE zips that start with 0
	if len(s) < 5 {
		z := "0"
		s := z + s

		return s
	}

	return s
}

func ConvertAllZips(r []*o.OrderRecord) []*o.OrderRecord {
	for _, v := range r {
		zipFiveDig := FirstFiveZip(v.PostalCode)
		v.PostalCode = zipFiveDig
	}
	return r
}

func CreateZipTable(r []*o.OrderRecord) []ZipTemp {
	records := ConvertAllZips(r)
	zipTempTable := []ZipTemp{}
	// zipTempUnit := ZipTemp{}
	for k, v := range records {
		z := ZipTemp{}
		z.Keys = append(z.Keys, k)
		z.Zip = v.PostalCode
		zipTempTable = append(zipTempTable, z)
		fmt.Println(z)
	}
	fmt.Printf("%T", zipTempTable)
	return zipTempTable
}

// TODO Sort Zip Table so ZipTemp contains a list of indexes per zip code
// There should only be 1 entry per Zip, with the list of indexes attached

//func SortZipTable(z []*ZipTemp) []ZipTemp {

//}
