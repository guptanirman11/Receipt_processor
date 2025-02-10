package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// In memory data structure
var (
	receiptStore = make(map[string]int)
	storeMutex   sync.Mutex
)

func ProcessReceipt(c *gin.Context) {

	var receipt Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid."})
		return
	}

	receiptID := uuid.New().String()

	points := CalculatePoints(receipt)

	storeMutex.Lock()
	receiptStore[receiptID] = points
	storeMutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"id": receiptID})

}

func isAlphanumeric(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsNumber(s)
}

func isMultipleOfQuarter(amount string) bool {
	// Convert the string to float
	num, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return false
	}
	// Check if the number is a multiple of 0.25
	return int(num*100)%25 == 0
}

func CalculatePoints(receipt Receipt) int {
	var points float64 = 0.00

	// One point for every alphanumeric character in the Retailer name.
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points += 1
		}
	}
	fmt.Println("DEBUG: Points after alphanumeric check", points)

	// 	50 points if the Total is a round dollar amount with no cents.
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 25 points if the Total is a multiple of 0.25.
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// 5 points for every two Items on the receipt.
	points += (math.Floor(float64(len(receipt.Items)) / 2.00)) * 5.00

	fmt.Println("DEBUG: Points after every two item check", points)

	// If the trimmed length of the item description is a multiple of 3, multiply the Price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {

		Price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			fmt.Printf("Skipping item: %s (Invalid Price: %s)\n", item.ShortDescription, item.Price)
			continue
		}

		trimmedLength := len(strings.TrimSpace(item.ShortDescription))

		if trimmedLength%3 == 0 {
			points += math.Ceil(float64(Price) * 0.2)
			fmt.Println("DEBUG: Points after item length check", points)

		}
	}

	// 6 points if the day in the purchase date is odd.
	if len(receipt.PurchaseDate) == 10 {
		dateParts := strings.Split(receipt.PurchaseDate, "-")
		if len(dateParts) == 3 {
			day, err := strconv.Atoi(dateParts[2])
			if err == nil && int(day)%2 == 1 {

				points += 6
			}
		}
	}
	fmt.Println("DEBUG: Points after purchase date odd check", points)
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if len(receipt.PurchaseTime) == 5 {
		purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)

		if err == nil {
			startTime, _ := time.Parse("15:04", "14:00") // 2:00 PM
			endTime, _ := time.Parse("15:04", "16:00")   // 4:00 PM

			if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
				points += 10
			}
		}

	}

	return int(points)

}

func GetPoints(c *gin.Context) {

	receiptID := c.Param("id")

	storeMutex.Lock()
	points, exist := receiptStore[receiptID]
	storeMutex.Unlock()

	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": int64(points)})

}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/receipts/process", ProcessReceipt)
	r.GET("/receipts/:id/points", GetPoints)

	r.Run(":8080") // Starting server on port 8080
}
