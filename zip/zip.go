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
	"sort"
	"time"
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
func FirstFiveZip(zip string) (string, error) {
	counter := 0
	for i := range zip {
		if i == 5 {
			return zip[:i], errors.New("couldn't truncate zip")
		}
		counter++
	}
	fmt.Printf("zip: %v\n", zip)

	return zip, nil
}

func convertAllZips(r []*o.OrderRecord) ([]*o.OrderRecord, error) {
	for _, v := range r {
		// Skips rows that are line items (empty fields)
		if (!isStringEmpty(v.BuyerFullName)) &&
			(!isStringEmpty(v.RecFullName)) {
			zipFiveDig, err := FirstFiveZip(v.PostalCode)
			if err != nil {
				fmt.Print("Zip code error ::", err)
			}
			fmt.Printf("zipFiveDig: %v\n", zipFiveDig)
			v.PostalCode = zipFiveDig
		}
		continue
	}
	return r, nil
}

func geocodeZips() ([][]string, error) {
	orderCsv, err := os.Open("./zip/GeoZip3.csv")
	if err != nil {
		fmt.Println("Couldn't Open GeoCode file! ::", err)
	}
	defer orderCsv.Close()

	geoZips, err := csv.NewReader(orderCsv).ReadAll()
	if err != nil {
		fmt.Println("Geocode Reader Error occured! ::", err)
	}

	return geoZips, nil
}

func findGeoCode(records [][]string, val string, col int) (GeoCode, error) {
	geoCode := GeoCode{}
	for _, row := range records {
		if row[col] == val {
			geoCode.Lat = row[1]
			geoCode.Lon = row[2]
		}
	}
	return geoCode, nil
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

func sleepAlert(t time.Duration) {
	time.Sleep(t * time.Millisecond)
	fmt.Println("Sleeping...")
}

func getWeatherData(orders []*o.OrderRecord, geoZips [][]string) ([]o.OrderRecord, error) {
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
		if !isStringEmpty(order.PostalCode) && !isStringEmpty(order.City) {
			sleepAlert(1100)
			gz, err := findGeoCode(geoZips, order.PostalCode, 0)
			if err != nil {
				fmt.Println("No GeoCode Found ::", err)
			}
			fmt.Printf("order.PostalCode: %v", order.PostalCode)
			temp, err := tempCheck(gz)
			if err != nil {
				fmt.Println("No Temp Found ::", err)
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
	return newOrders, nil
}

func CreateNewOrders(r []*o.OrderRecord) ([]o.OrderRecord, error) {
	orders, err := convertAllZips(r)
	if err != nil {
		fmt.Println("Zip conversion failure ::", err)
	}

	geoZips, err := geocodeZips()
	if err != nil {
		fmt.Println("Geocode Error occured! ::", err)
	}

	newOrders, err := getWeatherData(orders, geoZips)
	if err != nil {
		fmt.Printf("Coudn't get Weather Data ::")
	}

	return newOrders, nil
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

func tempCheck(gc GeoCode) (float64, error) {
	apiKey := o.GetKey()
	weather := o.WeatherData{}

	// manipulate lat/lon strings for api call
	lat := latitude(gc.Lat)
	lon := longitude(gc.Lon)

	// returns F instead of K
	imp := "&units=imperial"

	// 3 days of forecast instead of 5
	count := "&cnt=24"
	link := "https://api.openweathermap.org/data/2.5/forecast?"

	// parse the URL
	parsedUrl, err := url.Parse(link)
	if err != nil {
		fmt.Println("parsing error ::", err)
	}

	// make the call to the weather api
	resp, err := http.Get(parsedUrl.String() + lat + lon + apiKey + count + imp)
	if err != nil {
		fmt.Println("HTTP request error ::", err)
	}

	// read response data
	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read i/o error ::", err)
	}

	// unmarshal json into WeatherData format
	err = json.Unmarshal([]byte(respJSON), &weather)
	if err != nil {
		fmt.Printf("json unmarshalling error :: %v\n", err)
	}

	// find the max temp of the 24 received
	temp, err := findMaxTemp(weather.List)
	if err != nil {
		fmt.Printf("Couldn't find max temp :: %v\n", err)
	}

	// this dumb thing makes the float have 2 decimal for some reason
	return temp, nil
}

func findMaxTemp(r o.List) (float64, error) {
	// build an array of temps from the List
	var nums []float64
	var max float64
	for _, v := range r {
		nums = append(nums, v.Main.Temp_max)
	}

	//fmt.Printf("nums: %v \n", nums)

	sort.Float64s(nums)

	//fmt.Printf("nums2: %v \n", nums)

	if r != nil {
		max = nums[len(nums)-1]
	}

	//fmt.Printf("max: %v \n", math.Round(max))

	return math.Round(max), nil
}
