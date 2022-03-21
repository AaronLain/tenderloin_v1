package zip

import (
	o "ajl/tenderloin/orders"
	"fmt"
	"sort"
	"unicode/utf8"
)

type Keys []int

// The idea is to keep the keys for each order with their respective zip/temps
type ZipTemp struct {
	Zip      string
	Temp     string
	OrderNum []string
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
func FirstFiveZip(zip string) string {
	counter := 0
	for i := range zip {
		if i == 5 {
			return zip[:i]
		}
		counter++
	}
	// Adds a zero to NE zips that start with 0
	if len(zip) < 5 {
		z := "0"
		s := z + zip

		return s
	}
	return zip
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

func containsKey(a, b Keys) bool {
	sort.Ints(a)
	sort.Ints(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TODO Sort Zip Table so ZipTemp contains a list of indexes per zip code
// There should only be 1 entry per Zip, with the list of indexes attached

// []ZipTemp
func SortZipTable(z []ZipTemp) {
	m := make(map[string][]ZipTemp)
	for _, o := range z {
		m[o.Zip] = append(m[o.Zip], o)
		fmt.Printf("o: %v \n", o)
	}
	fmt.Printf("newZipTable: %v \n", m)

}

//  zipTable := z
// newZipTable := []ZipTemp{}
// lastRow := len(zipTable) - 1
// for i, v := range zipTable {
// 	//fmt.Printf("i: %v\n ", i)
// 	thisRow := lastRow - i
// 	thisZip := v.Zip
// 	theseKeys := v.Keys
// 	ztemp := ZipTemp{}
// 	if thisRow <= lastRow {
// 		for _, vv := range zipTable {
// 			noDupeKeys := !containsKey(vv.Keys, theseKeys)
// 			// Need to make sure no duplicate keys are added (probably refactor later)
// 			if thisZip == vv.Zip {
// 				if noDupeKeys {
// 					theseKeys = append(theseKeys, vv.Keys...)
// 				}

// 				ztemp.Keys = theseKeys

// 				if ztemp.Zip != thisZip {
// 					ztemp.Zip = thisZip
// 					newZipTable = append(newZipTable, ztemp)
// 				}
// 			}
// 		}
// 	}
// }

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

//[]ZipTemp
func CreateZipTable(r []*o.OrderRecord) {
	records := ConvertAllZips(r)

	zipTempTable := []ZipTemp{}
	// zipTempUnit := ZipTemp{}
	for _, v := range records {
		z := ZipTemp{}
		//TODO ADD [] of ORDER NUMBERS FOR VERIFICATION
		if !isStringEmpty(v.PostalCode) {
			z.Zip = v.PostalCode
			// Make sure we trim the octothorpe from Order Number
			z.OrderNum = append(z.OrderNum, trimFirstRune(v.OrderNum))
			zipTempTable = append(zipTempTable, z)
		}
		continue
	}

	//fmt.Printf("%T", zipTempTable)
	SortZipTable(zipTempTable)
}
