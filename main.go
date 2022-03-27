package main

import (
	orders "ajl/tenderloin/orders"
	zip "ajl/tenderloin/zip"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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

func csvWriter(input string, o []*orders.OrderRecord) {
	// string manipulation because stuff is picky
	output1 := strings.TrimSuffix(input, ".csv")
	output2 := strings.TrimPrefix(output1, "./")
	outputName := output2 + "_"
	// get the temps and bring back the fresh data
	newRecords, err := zip.GetTemps(o)
	// random number for filename creation
	randomTime := rand.NewSource(time.Now().UnixNano())
	randSuffix := randomTime.Int63()
	str_randSuffix := strconv.FormatInt(randSuffix, 10)
	if err != nil {
		fmt.Println("Failed to get new records ::", err)
	}
	// check to see if filename already exists before creating
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		//this puts the csv in the local file
		file, err := os.Create(outputName + str_randSuffix + ".csv")
		if err != nil {
			fmt.Println("Can't create csv ::", err)
		}
		gocsv.MarshalFile(&newRecords, file)
	}
}

func initializeCSV() {
	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	records, err := csvReader(fileName)
	if err != nil {
		fmt.Println("Can't initialize reader ::", err)
	}

	csvWriter(fileName, records)

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
