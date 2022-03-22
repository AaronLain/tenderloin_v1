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

type LineItem struct {
	OrderNum      string `csv:"Name"`
	ItemSKU       string `csv:"Lineitem sku"`
	ItemName      string `csv:"Lineitem name"`
	ItemUnitPrice string `csv:"Lineitem price"`
}

type Coordinates struct {
	Lat int
	Lon int
}

type Weather []struct {
	ID          int
	weatherMain string
	Description string
	Icon        string
}

type Main struct {
	Temp      float32
	FeelsLike float32
	TempMin   float32
	TempMax   float32
	Pressure  int
	Humidity  int
}

type Clouds struct {
	All int
}

type Sys struct {
	Type    int
	ID      int
	Country string
	Sunrise int
	Sunset  int
}

type Wind struct {
	Speed float32
	Deg   int
	Gust  float32
}

type WeatherData struct {
	Coordinates Coordinates
	Weather     Weather
	Base        string
	Main        Main
	Visibility  int
	Wind        Wind
	Clouds      Clouds
	Dt          int
	Sys         Sys
	Timezone    int
	ID          int
	Name        string
	Cod         int
}

// []*OrderRecord
// func AddLineItems(records []*OrderRecord) {
// 	// newRecords := []OrderRecord{}
// 	lastRow := records[len(records)-1]

// 	for i, v := range records {
// 		order := v

// 		fmt.Printf("i: %d\n", i)
// 		fmt.Printf("v: %v\n", v)

// 		// Runs until the last row
// 		if lastRow != nil {
// 			if (isStringEmpty(order.BuyerFullName)) && (isStringEmpty(order.RecFullName)) {
// 				thisOrder := order
// 				nextOrder := records[i+1]
// 				prevOrder := records[i-1]
// 				lineItem := []LineItem{}
// 				// if the next order has these empty fields, it is a line item and should be added to the order
// 				if (isStringEmpty(nextOrder.BuyerFullName)) && (isStringEmpty(nextOrder.RecFullName)) {
// 					thisOrder.LineItem = append(nextOrder.LineItems)
// 				}

// 			}
// 		}

// 	}
// 	fmt.Printf("Last Row %T\n", lastRow)
// }
