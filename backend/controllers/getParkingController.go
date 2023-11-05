package controllers

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetParking(c *gin.Context) {
	// URL to forward the GET request to
	targetURL := "https://hackathon.kojikukac.com/api/ParkingSpot/getAll"
	token := "a70b5e43-4881-45e5-8a6b-858628c93ab6"

	// Create a new GET request to the target URL
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could not create forward request",
		})
		c.Abort()
		return
	}

	// Add headers to the request, including the token
	req.Header.Set("Api-Key", token)

	// Perform the GET request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could not get forward request",
		})
		c.Abort()
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could not read response body",
		})
		c.Abort()
		return
	}

	err = json.Unmarshal(body, &models.ParkingSpots)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(models.ParkingSpots)
	// Set the status code of the forwarded response and send the response
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
