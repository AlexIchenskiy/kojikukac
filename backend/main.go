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
	database.GlobalDB.AutoMigrate(&models.User{})

	// Initialize Router
	router := gin.Default()
	api := router.Group("/api")
	user := api.Group("/user").Use(middleware.Cors())
	user.POST("/login", controllers.Login)
	user.POST("/register", controllers.Register).Use(middleware.Cors())
	user.GET("/", controllers.Profile).Use(middleware.Auth(), middleware.Cors())

	reservation := api.Group("/reservation").Use(middleware.Cors()).Use(middleware.Cors())
	reservation.POST("/delete", controllers.Profile).Use(middleware.Auth())
	reservation.POST("/add", controllers.AddReservation).Use(middleware.Auth())

	api.POST("/search").Use(middleware.Cors())

	router.Run(":8080")
}
