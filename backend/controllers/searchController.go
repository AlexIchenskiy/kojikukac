package controllers

import (
	"backend/models"
	"math"

	"github.com/gin-gonic/gin"
)

type SearchType struct {
	Type      string  `json:"type" binding:"required"`
	Value     float64 `json:value`
	Latitude  float64 `json:"value" binding:"required"`
	Longitude float64 `json:"value" binding:"required"`
}

func Search(c *gin.Context) {

	var t SearchType

	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	var returnVal []models.ParkingSpot

	if t.Type == "location" {
		for _, i := range models.ParkingSpots {
			distance := math.Sqrt(math.Pow((i.Latitude-t.Latitude), 2) + math.Pow((i.Longitude-float64(t.Longitude)), 2))

			if distance < t.Value {
				returnVal = append(returnVal, i)
			}

		}
	} else if t.Type == "spottype" {
		/*for _,i := models.ParkingSpots {


		}
		*/
	} else if t.Type == "price" {
		for _, i := range models.ParkingSpots {

			if i.Price < t.Value {
				returnVal = append(returnVal, i)
			}

		}

	} else {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}

}
