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

// Apparently this is how we have to do things
// but at least it (probably) works!
func isStringEmpty(str ...string) bool {
	for _, s := range str {
		if s == "" {
			return true
		}
	}
	return false
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
		// Skips rows that are line items (empty fields)
		if (!isStringEmpty(v.BuyerFullName)) && (!isStringEmpty(v.RecFullName)) {
			zipFiveDig := FirstFiveZip(v.PostalCode)
			v.PostalCode = zipFiveDig
		}
		continue
	}
	return r
}

// TODO Sort Zip Table so ZipTemp contains a list of indexes per zip code
// There should only be 1 entry per Zip, with the list of indexes attached

//func SortZipTable(z []*ZipTemp) []ZipTemp {

//}

//[]ZipTemp
func CreateZipTable(r []*o.OrderRecord) {
	records := ConvertAllZips(r)
	zipTempTable := []ZipTemp{}
	// zipTempUnit := ZipTemp{}
	for i, v := range records {
		z := ZipTemp{}
		z.Keys = append(z.Keys, i)
		z.Zip = v.PostalCode
		zipTempTable = append(zipTempTable, z)
		fmt.Println(z)
	}
	fmt.Printf("%T", zipTempTable)
	// return zipTempTable
}
