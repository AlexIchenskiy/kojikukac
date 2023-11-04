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
	user := api.Group("/user")
	user.POST("/user/login", controllers.Login)
	user.POST("/user/register", controllers.Register)
	user.GET("/user", controllers.Profile).Use(middleware.Auth())

	reservation := api.Group("/reservation")
	reservation.POST("/add", controllers.Profile).Use(middleware.Auth())
	reservation.POST("/delite", controllers.Profile).Use(middleware.Auth())

	search := api.Group("/search")
	search.POST("/location")
	search.POST("/type")
	search.POST("/availability")
	search.POST("/price")

	router.Run(":8080")
}
