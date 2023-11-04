package controllers

import (
	"github.com/gin-gonic/gin"
)

type Resevation struct {
	Email         string `json:"email" binding:"required"`
	ParkingSpotId string `json:"parkingspotid" binding:"required"`
	EndH          string `json:"endh" binding:"required"`
	EndM          string `json:"endm" binding:"required"`
}

func addRegistration(c *gin.Context) {
	/*// HTTP endpoint
	posturl := "https://jsonplaceholder.typicode.com/posts"

	// JSON body
	body := []byte(`{
		"title": "Post title",
		"body": "Post description",
		"userId": 1
	}`)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	*/
}
