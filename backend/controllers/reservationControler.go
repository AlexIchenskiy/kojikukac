package controllers

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddReservation(c *gin.Context) {
	// HTTP endpoint
	postURL := "https://hackathon.kojikukac.com/swagger/api/ParkingSpot/reserve"
	token := "a70b5e43-4881-45e5-8a6b-858628c93ab6"

	var reservation models.Resevation
	err := c.ShouldBindJSON(&reservation)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs ",
		})
		c.Abort()
		return
	}

	// JSON body
	data := map[string]interface{}{
		"parkingSpotId": reservation.ParkingSpotId,
		"endH":          reservation.EndH,
		"endM":          reservation.EndM,
	}

	// Convert the map to a JSON string
	jsonString, err := json.Marshal(data)
	if err != nil {
        log.Fatalf("%v", err)
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
		c.Abort()
		return
	}
	// Create a HTTP post request
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonString))
	if err != nil {
		panic(err)
	}

	// Add the token to the request headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check HTTP response status code
	if resp.StatusCode != http.StatusOK {
		c.JSON(500, gin.H{
			"Error": "Error Creating Reservation",
		})
		c.Abort()
		return
	}

	err1 := reservation.CreateReservation()
	if err1 != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating Reservation",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully reserved",
	})

}
