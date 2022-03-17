package zip

import (
	o "ajl/tenderloin/orders"
	"fmt"
)

type Keys []int

// The idea is to keep the keys for each order with their respective zip/temps
type ZipTemp struct {
	Id int
	Keys
	Zip  string
	Temp string
}

type ZipTempBool interface {
	exists()
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

//
// func compareZipTemps(a, b ZipTemp) bool {

// }

func ConvertAllZips(r []*o.OrderRecord) []*o.OrderRecord {
	for _, v := range r {
		//fmt.Printf("records: %d %v\n", i, v.PostalCode)
		// Skips rows that are line items (empty fields)
		if (!isStringEmpty(v.BuyerFullName)) && (!isStringEmpty(v.RecFullName)) {
			zipFiveDig := FirstFiveZip(v.PostalCode)
			v.PostalCode = zipFiveDig
		}
		continue
	}
	return r
}

func containsKey(z []int, c []int) bool {
	if z == c {
		return true
	}
	if a.X != b.X || a.Y != b.Y {
		return false
	}
	if len(a.Z) != len(b.Z) || len(a.M) != len(b.M) {
		return false
	}
	for i, v := range a.Z {
		if b.Z[i] != v {
			return false
		}
	}
	for k, v := range a.M {
		if b.M[k] != v {
			return false
		}
	}
	return true
}

// TODO Sort Zip Table so ZipTemp contains a list of indexes per zip code
// There should only be 1 entry per Zip, with the list of indexes attached

// []ZipTemp
func SortZipTable(z []ZipTemp) {
	zipTable := z
	newZipTable := []ZipTemp{}
	lastRow := len(zipTable) - 1
	fmt.Printf("last row: %v \n", lastRow)
	for i, v := range zipTable {
		thisRow := len(zipTable) - i
		fmt.Printf("this row: %v \n", thisRow)
		thisZip := v.Zip
		theseKeys := v.Keys
		ztemp := ZipTemp{}
		if thisRow <= lastRow {
			for _, vv := range zipTable {
				// Need to make sure no duplicate keys are added (probably refactor later)
				if thisZip == vv.Zip {
					if !containsKey(theseKeys, vv.Keys) {
						theseKeys = append(theseKeys, vv.Keys...)
					}

					ztemp.Keys = theseKeys

					if ztemp.Zip != thisZip {
						ztemp.Zip = thisZip
						newZipTable = append(newZipTable, ztemp)
					}

				}
			}
		}
	}
	fmt.Printf("newZipTable: %v \n", newZipTable)

}

// fmt.Printf("ZipTable: %v", zipTable)

//[]ZipTemp
func CreateZipTable(r []*o.OrderRecord) {
	records := ConvertAllZips(r)

	zipTempTable := []ZipTemp{}
	// zipTempUnit := ZipTemp{}
	for i, v := range records {
		z := ZipTemp{}
		//TODO ADD [] of ORDER NUMBERS FOR VERIFICATION
		if !isStringEmpty(v.PostalCode) {
			z.Keys = append(z.Keys, i)
			z.Zip = v.PostalCode
			zipTempTable = append(zipTempTable, z)
		}
		continue
	}
	//fmt.Printf("%T", zipTempTable)
	SortZipTable(zipTempTable)
}
