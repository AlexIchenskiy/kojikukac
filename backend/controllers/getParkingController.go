package controllers

import (
	"backend/models"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ZoneAPriceMax float64 = 1.0
var ZoneBPriceMax float64 = 4.0
var ZoneCPriceMax float64 = 0.5
var ZoneDPriceMax float64 = 3.0
var ZoneAPriceMin float64 = 0.15
var ZoneBPriceMin float64 = 1.5
var ZoneCPriceMin float64 = 0.01
var ZoneDPriceMin float64 = 0.1

func GetParking(c *gin.Context) {
	// URL to forward the GET request to
	targetURL := "https://hackathon.kojikukac.com/api/ParkingSpot/getAll"
	token := "a70b5e43-4881-45e5-8a6b-858628c93ab6"

	// Create a new GET request to the target URL
	req, err := http.NewRequest("GET", targetURL, nil)
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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %s", err)
		c.JSON(500, gin.H{
			"Error": "Could not read response body",
		})
		c.Abort()
		return
	}

	err = json.Unmarshal(body, &models.ParkingSpots)
	if err != nil {
		// at least an hour has passed trying to figure out what's wrong here, only to find out
		// someone used log.Fatalf(). Remember, log.Fatalf() will completely kill the backend so
		// DO NOT JUST USE IT WILLY NILLY, AND ONLY USE IT WHERE THE BACKEND IS EXPECTED TO NOT
		// WORK PROPERLY, USE log.Printf() EVERYWHERE ELSE INSTEAD!
		log.Printf("Error: %s", err)
	}
	counterOccupiedA := 0
	counterTotalA := 0
	counterOccupiedB := 0
	counterTotalB := 0
	counterOccupiedC := 0
	counterTotalC := 0
	counterOccupiedD := 0
	counterTotalD := 0
	for _, i := range models.ParkingSpots {
		if i.ParkingSpotZone == "Zone1" {
			if i.Occupied {
				counterOccupiedA++
			}
			counterTotalA++
		}
		if i.ParkingSpotZone == "Zone2" {
			if i.Occupied {
				counterOccupiedB++
			}
			counterTotalB++
		}
		if i.ParkingSpotZone == "Zone3" {
			if i.Occupied {
				counterOccupiedC++
			}
			counterTotalC++
		}
		if i.ParkingSpotZone == "Zone4" {
			if i.Occupied {
				counterOccupiedD++
			}
			counterTotalD++
		}
	}

	var koefA float64 = (ZoneAPriceMax - ZoneAPriceMin) / float64(counterTotalA-counterOccupiedA)
	var koefB float64 = (ZoneBPriceMax - ZoneBPriceMin) / float64(counterTotalB-counterOccupiedB)
	var koefC float64 = (ZoneCPriceMax - ZoneCPriceMin) / float64(counterTotalC-counterOccupiedC)
	var koefD float64 = (ZoneDPriceMax - ZoneDPriceMin) / float64(counterTotalD-counterOccupiedD)

	for _, i := range models.ParkingSpots {
		if i.ParkingSpotZone == "Zone1" {
			i.Price = ZoneAPriceMin + float64(counterOccupiedA)*koefA
		}
		if i.ParkingSpotZone == "Zone2" {
			i.Price = ZoneBPriceMin + float64(counterOccupiedB)*koefB
		}
		if i.ParkingSpotZone == "Zone3" {
			i.Price = ZoneCPriceMin + float64(counterOccupiedC)*koefC
		}
		if i.ParkingSpotZone == "Zone4" {
			i.Price = ZoneDPriceMin + float64(counterOccupiedD)*koefD
		}

	}

	// Set the status code of the forwarded response and send the response
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
