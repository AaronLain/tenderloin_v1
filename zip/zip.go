package zip

import (
	o "ajl/tenderloin/orders"
	"encoding/csv"
	"fmt"
	"os"
	"unicode/utf8"
)

type OrderNum []string

// The idea is to keep the keys for each order with their respective zip/temps
type ZipTemp struct {
	Zip  string
	Temp string
	OrderNum
}

type GeoZip struct {
	ZipCode string `csv:"ZIP"`
	Lat     string `csv:"LAT"`
	Lon     string `csv:"LNG"`
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

func geocodeZips() [][]string {
	orderCsv, err := os.Open("./zip/ZipGeoCode.csv")
	if err != nil {
		fmt.Println("Error occured! ::", err)
	}
	defer orderCsv.Close()

	geoZips, err := csv.NewReader(orderCsv).ReadAll()
	if err != nil {
		panic(err)
	}

	return geoZips
}

// func containsKey(a, b Keys) bool {
// 	sort.Ints(a)
// 	sort.Ints(b)

// 	for i, v := range a {
// 		if v != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

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

// remove octothorpe (first rune) from OrderNum
func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

//[]GeoZip
func find(records [][]string, val string, col int) {
	for _, row := range records {
		if row[col] == val {
			fmt.Printf("row: %v \n", row)
		}
		continue
	}

}

// func find(records [][]string, val string, col int) {
// 	//newGZ := GeoZip{}
// 	for _, row := range records {
// 		if row[col] == val {
// 			fmt.Printf("row: %v", row)
// 		}
// 	}

// }

//[]*o.OrderRecord
func GetTemps(r []*o.OrderRecord) {
	orders := ConvertAllZips(r)
	geoZips := geocodeZips()
	newOrders := []o.OrderRecord{}

	// zipTempTable := []ZipTemp{}
	// zipTempUnit := ZipTemp{}
	for _, order := range orders {
		iceProfile := "0"
		if (!isStringEmpty(order.BuyerFullName)) && (!isStringEmpty(order.RecFullName)) {
			thisOrder := o.OrderRecord{
				OrderNum:        order.OrderNum,
				OrderDate:       order.OrderDate,
				DatePaid:        order.DatePaid,
				Total:           order.Total,
				AmountPaid:      order.AmountPaid,
				Tax:             order.Tax,
				ShippingPaid:    order.ShippingPaid,
				ShippingService: order.ShippingService,
				CustomField1:    order.CustomField1,
				CustomField2:    order.CustomField2,
				CustomField3:    iceProfile,
				Source:          order.Source,
				BuyerFullName:   order.BuyerFullName,
				BuyerEmail:      order.BuyerEmail,
				BuyerPhone:      order.BuyerPhone,
				RecPhone:        order.RecPhone,
				City:            order.City,
				State:           order.State,
				PostalCode:      order.PostalCode,
				ItemSKU:         order.ItemSKU,
				ItemUnitPrice:   order.ItemUnitPrice,
				ItemName:        order.ItemName,
			}

			find(geoZips, order.PostalCode, 0)
			thisOrder.PostalCode = "touched by an angel (func)"
			// row := find(geoZips, v.PostalCode, 0)
			newOrders = append(newOrders, thisOrder)
			// fmti.Printf("row %v \n", row)
			//fmt.Printf("orderFullName: %v \n", thisOrder.BuyerFullName)
		}

	}

	fmt.Printf("newRecords: %v \n", newOrders)

}
