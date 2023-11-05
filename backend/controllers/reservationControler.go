package controllers

import (
	"backend/database"
	"backend/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddReservationPayload struct {
    parkingSpotId int
    endH int
    endM int
}

func AddReservation(c *gin.Context) {
    // Initialize a user model
    var user models.User
    // Get the email from the authorization middleware
    email, exists := c.Get("email") 
    if !exists {
        c.JSON(404, gin.H{
            "Error": "Token not provided",
        })
        c.Abort()
        return
    }
    if email, ok := email.(string); ok {
        // Query the database for the user
        result := database.GlobalDB.Where("email = ?", email).First(&user)
        // If the user is not found, return a 404 status code
        if result.Error == gorm.ErrRecordNotFound {
            c.JSON(404, gin.H{
                "Error": "User Not Found",
            })
            c.Abort()
            return
        }
        // If an error occurs while retrieving the user profile, return a 500 status code
        if result.Error != nil {
            c.JSON(500, gin.H{
                "Error": "Could Not Get User Profile",
            })
            c.Abort()
            return
        }

        var payload AddReservationPayload
        err := c.ShouldBindJSON(&payload)
        if err != nil {
            c.JSON(400, gin.H{
                "Error": "Bad Request",
            })
            c.Abort()
            return
        }

        // HTTP endpoint
        postURL := "https://hackathon.kojikukac.com/swagger/api/ParkingSpot/reserve"
        token := "a70b5e43-4881-45e5-8a6b-858628c93ab6"

        req, err := http.NewRequest("POST", postURL, nil)
        if err != nil {
            log.Printf("Error: %s", err)
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
            log.Printf("Error: %s", err)
            c.JSON(500, gin.H{
                "Error": "Could not get forward request",
            })
            c.Abort()
            return
        }
        defer resp.Body.Close()

        reservation := models.Resevation {};

        // JSON body
        data := map[string]interface{}{
            "parkingSpotId": payload.parkingSpotId,
            "endH": payload.endH,
            "endM": payload.endM,
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
        req, err = http.NewRequest("POST", postURL, bytes.NewBuffer(jsonString))
        if err != nil {
            panic(err)
        }

        // Add the token to the request headers
        req.Header.Set("Authorization", "Bearer "+token)
        req.Header.Set("Content-Type", "application/json")

        // Make the HTTP request
        client := &http.Client{}
        resp, err = client.Do(req)
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

        email, _ := c.Get("email")
        reservation.Email = email.(string)
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
}
