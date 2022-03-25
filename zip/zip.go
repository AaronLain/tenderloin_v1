package zip

import (
	o "ajl/tenderloin/orders"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
)

type OrderNum []string

// The idea is to keep the keys for each order with their respective zip/temps
type ZipTemp struct {
	Zip  string
	Temp string
	OrderNum
}

type GeoCode struct {
	Lat string `csv:"LAT"`
	Lon string `csv:"LNG"`
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

// Keeps only base zip
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

func ConvertAllZips(r []*o.OrderRecord) ([]*o.OrderRecord, error) {
	for _, v := range r {
		// Skips rows that are line items (empty fields)
		if (!isStringEmpty(v.BuyerFullName)) && (!isStringEmpty(v.RecFullName)) {
			zipFiveDig := FirstFiveZip(v.PostalCode)
			v.PostalCode = zipFiveDig
		}
		continue
	}
	return r, errors.New("Couldn't convert zip codes")
}

func geocodeZips() ([][]string, error) {
	orderCsv, err := os.Open("./zip/ZipGeoCode.csv")
	if err != nil {
		fmt.Println("Couldn't Open GeoCode file! ::", err)
	}
	defer orderCsv.Close()

	geoZips, err := csv.NewReader(orderCsv).ReadAll()
	if err != nil {
		fmt.Println("Geocode Reader Error occured! ::", err)
	}

	return geoZips, err
}

func findGeoCode(records [][]string, val string, col int) (GeoCode, error) {
	geoCode := GeoCode{}
	for _, row := range records {
		if row[col] == val {
			geoCode.Lat = row[1]
			geoCode.Lon = row[2]
		}
	}
	return geoCode, errors.New("Couldn't find GeoCode")
}

func profileAssignment(temp float64) string {
	if temp <= 55.0 {
		return "Profile 1"
	} else if (temp > 55.0) && (temp <= 75.0) {
		return "Profile 2"
	} else if (temp > 75) && (temp <= 85) {
		return "Profile 3"
	} else if (temp > 85) && (temp <= 95) {
		return "Profile 4"
	} else if temp > 95 {
		return "Profile 5"
	} else {
		return "No Temp Found"
	}

}

//[]*o.OrderRecord
func GetTemps(r []*o.OrderRecord) ([]o.OrderRecord, error) {
	orders, err := ConvertAllZips(r)
	if err != nil {
		fmt.Println("Zip conversion failure ::", err)
	}
	geoZips, err := geocodeZips()
	if err != nil {
		fmt.Println("Geocode Error occured! ::", err)
	}

	newOrders := []o.OrderRecord{}

	for _, order := range orders {
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
			CustomField3:    order.CustomField3,
			AvgTemp:         order.AvgTemp,
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
		// Where the magic happens
		// find the geocode, check the temp, apply the data accordingly
		if !isStringEmpty(order.PostalCode) {
			gz, err := findGeoCode(geoZips, order.PostalCode, 0)
			if err != nil {
				fmt.Println("Find GeoCode Failed ::", err)
			}
			temp, err := tempCheck(gz)
			if err != nil {
				fmt.Println("Temperature Check Failed ::", err)
			}
			thisOrder.AvgTemp = temp
			thisOrder.CustomField3 = profileAssignment(temp)
			newOrders = append(newOrders, thisOrder)
		} else if isStringEmpty(order.PostalCode) {
			newOrders = append(newOrders, thisOrder)
		} else {
			fmt.Println("Get Temps Failed")
		}

	}

	return newOrders, err
}

func longitude(input string) string {
	lon1 := "lon="
	lon2 := input
	lon3 := "&"
	output := lon1 + lon2 + lon3

	return output
}

func latitude(input string) string {
	lat1 := "lat="
	lat2 := input
	lat3 := "&"
	output := lat1 + lat2 + lat3

	return output
}

// string
func tempCheck(gc GeoCode) (float64, error) {
	apiKey := o.GetKey()
	weather := o.WeatherData{}
	// TODO This needs to run at 60 req/minute
	// whether that's a quick burst of 60 and a pause
	// or one req every ~1.1 seconds
	lat := latitude(gc.Lat)
	lon := longitude(gc.Lon)
	// returns F instead of K
	imp := "&units=imperial"
	// 3 days of forecast instead of 5
	count := "&cnt=24"
	link := "https://api.openweathermap.org/data/2.5/forecast?"

	parsedUrl, err := url.Parse(link)
	if err != nil {
		fmt.Println("parsing error ::", err)
	}

	resp, err := http.Get(parsedUrl.String() + lat + lon + apiKey + count + imp)
	if err != nil {
		fmt.Println("HTTP request error ::", err)
	}

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read i/o error ::", err)
	}

	err = json.Unmarshal([]byte(respJSON), &weather)
	if err != nil {
		fmt.Println("json unmarshalling error ::", err)
	}
	temp, err := tempAvg(weather.List)
	// fmt.Printf("avg: %v \n", temp)
	return (math.Round(temp*100) / 100), err
}

func tempAvg(r o.List) (float64, error) {
	total := 0.0
	len := float64(len(r))
	for _, v := range r {
		total = total + v.Main.Temp
		// fmt.Printf("dt: %v \n", v.Dt_txt)
	}

	avg := total / len

	return avg, errors.New("couldn't find average temperature")

}
