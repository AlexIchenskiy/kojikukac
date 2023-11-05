package main

import (
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
	database.GlobalDB.AutoMigrate(&models.User{}, &models.Session{})

	// Initialize Router
	router := gin.Default()
	router.Use(middleware.Cors(), middleware.Auth())
	api := router.Group("/api")
	user := api.Group("/user")
	user.POST("/login", controllers.Login).Use(middleware.Cors())
	user.POST("/register", controllers.Register).Use(middleware.Cors())
	user.GET("/profile", controllers.Profile).Use(middleware.Cors(), middleware.Auth())

	reservation := api.Group("/reservation")
	reservation.POST("/add", controllers.AddReservation).Use(middleware.Auth(), middleware.Cors())
	// reservation.DELETE("/delete", controllers.AddReservation).Use(middleware.Auth(), middleware.Cors())
	// reservation.DELETE("/delete", controllers.AddReservation).Use(middleware.Auth(), middleware.Cors())

	api.GET("/search", controllers.Search).Use(middleware.Cors())

	parkingSpot := api.Group("/ParkingSpot")
	parkingSpot.GET("/getAll", controllers.GetParking).Use(middleware.Cors(), middleware.Auth())
	parkingSpot.POST("/", controllers.GetParking).Use(middleware.Cors(), middleware.Auth())
	parkingSpot.DELETE("/{id}", controllers.GetParking).Use(middleware.Cors(), middleware.Auth())

	//go consumer.Consume()

	router.Run(":8080")
}
