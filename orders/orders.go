package orders

type OrderRecord struct {
	OrderNum        string `csv:"Name"`
	OrderDate       string `csv:"Created at"`
	DatePaid        string `csv:"Paid at"`
	Total           string `csv:"Total"`
	AmountPaid      string `csv:"Total"`
	Tax             string `csv:"Taxes"`
	ShippingPaid    string `csv:"Shipping"`
	ShippingService string `csv:"Shipping Method"`
	CustomField1    string `csv:"-"`
	CustomField2    string `csv:"Tags"`
	CustomField3    string `csv:"-"`
	Source          string `csv:"Source"`
	BuyerFullName   string `csv:"Billing Name"`
	BuyerEmail      string `csv:"Email"`
	BuyerPhone      string `csv:"Billing Phone"`
	RecFullName     string `csv:"Shipping Name"`
	RecPhone        string `csv:"Shipping Phone"`
	RecCompany      string `csv:"Shipping Company"`
	AddressLine1    string `csv:"Shipping Address1"`
	AddressLine2    string `csv:"Shipping Address2"`
	City            string `csv:"Shipping City"`
	State           string `csv:"Shipping Province"`
	PostalCode      string `csv:"Shipping Zip"`
	CountryCode     string `csv:"Shipping Country"`
	ItemSKU         string `csv:"Lineitem sku"`
	ItemName        string `csv:"Lineitem name"`
	ItemUnitPrice   string `csv:"Lineitem price"`
}

// Weather Data Structure
type Coord struct {
	Lat float64
	Lon float64
}

type Weather []struct {
	ID          int
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temp       float64
	Feels_like float64
	Temp_min   float64
	Temp_max   float64
	Pressure   int32
	Sea_level  int32
	Grnd_level int32
	Humidity   int32
	Temp_kf    float64
}

type Clouds struct {
	All int
}

type Sys struct {
	Pod string
}

type Wind struct {
	Speed float64
	Deg   int
	Gust  float64
}

type City struct {
	ID         int32
	Name       string
	Coord      Coord
	Country    string
	Population int32
	Timezone   int
	Sunrise    int32
	Sunset     int32
}

type List []struct {
	Dt         int32
	Main       Main
	Weather    Weather
	Clouds     Clouds
	Wind       Wind
	Visibility int
	Pop        float64
	Sys        Sys
	Dt_txt     string
}

type WeatherData struct {
	Cod     string
	Message int
	Cnt     int
	List    List
	City    City
}
