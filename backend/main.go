package main

import (
	"backend/consumer"
	"backend/controllers"
	"backend/database"
	"backend/middleware"
	"backend/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	err := database.InitDatabase()
	if err != nil {
		// Log the error and exit
		log.Fatalln("could not create database", err)
	}
	// Automigrate the User model
	// AutoMigrate() automatically migrates our schema, to keep our schema upto date.
	database.GlobalDB.AutoMigrate(&models.User{})

	// Initialize Router
	router := gin.Default()
	router.Use(middleware.Cors())
	api := router.Group("/api")
	user := api.Group("/user")
	user.POST("/login", controllers.Login)
	user.POST("/register", controllers.Register)
	user.GET("/", controllers.Profile).Use(middleware.Auth())

	reservation := api.Group("/reservation")
	reservation.POST("/add", controllers.AddReservation).Use(middleware.Auth(), middleware.Cors())

	api.GET("/ParkingSpot/getAll", controllers.GetParking).Use(middleware.Auth(), middleware.Cors())
	api.GET("/ParkingSpot/update", controllers.GetParking).Use(middleware.Auth(), middleware.Cors())

	api.POST("/search").Use(middleware.Cors())

	consumer.Consume()

	router.Run(":8080")
}
