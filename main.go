package main

import (
	orders "ajl/tenderloin/orders"
	zip "ajl/tenderloin/zip"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func csvReader(s string) ([]*orders.OrderRecord, error) {
	recordFile, err := os.Open(s)
	if err != nil {
		fmt.Println("Reader Error occured! ::", err)
	}

	records := []*orders.OrderRecord{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		fmt.Println("Unmarshalling Error occured! ::", err)
	}
	defer recordFile.Close()

	return records, err
}

func csvWriter(input string, days int, o []*orders.OrderRecord) {
	// string manipulation because stuff is picky
	output1 := strings.TrimSuffix(input, ".csv")
	outputName := strings.TrimPrefix(output1, "./")

	// get the temps and bring back the fresh data
	newRecords, err := zip.CreateNewOrders(o, days)
	if err != nil {
		fmt.Println("Failed to get new records ::", err)
	}

	// check to see if filename already exists before creating
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		//this puts the csv in the local file
		file, err := os.Create(outputName + ".csv")
		if err != nil {
			fmt.Println("Can't create csv ::", err)
		}
		gocsv.MarshalFile(&newRecords, file)
	}
}

func initializeCSV() {
	localString := "./"
	inputFile := strings.Join(os.Args[2:], "")
	fileName := localString + inputFile

	records, err := csvReader(fileName)
	if err != nil {
		fmt.Println("Can't initialize reader ::", err)
	}

	daysPtr := flag.Int("days", 4, "days for forcast")

	flag.Parse()

	days := *daysPtr

	fmt.Println("INITIALIZE Days: ", days)

	now := time.Now().Format("20060102150405")

	outputFileName := "Output_" + now

	csvWriter(outputFileName, days, records)

}

func initializeRouter() {
	router := gin.Default()

	router.Use(static.Serve("/api", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	api.GET("/zip", zipHandler)

	router.Run(":8080")
}

func zipHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "zipHandler not yet implemented",
	})
}

func main() {

	//initializeRouter()

	initializeCSV()
}
