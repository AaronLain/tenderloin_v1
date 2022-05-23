package orders

type OrderRecord struct {
	OrderNum      string  `csv:"Order - Number"`
	CustomField3  string  `csv:"CustomField3"`
	AvgTemp       float64 `csv:"AvgTemp"`
	City          string  `csv:"Ship To - City"`
	State         string  `csv:"Ship To - State"`
	PostalCode    string  `csv:"Ship To - Postal Code"`
	CountryCode   string  `csv:"Shipping Country"`
	ItemSKU       string  `csv:"Lineitem sku"`
	ItemName      string  `csv:"Lineitem name"`
	ItemUnitPrice string  `csv:"Lineitem price"`
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
