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
	router := setupRouter()
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	// Create a new router
	r := gin.Default()
	// Add a welcome route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	// Create a new group for the API
	api := r.Group("/api")
	{
		// Create a new group for the public routes
		public := api.Group("/public")
		{
			// Add the login route
			public.POST("/login", controllers.Login)
			// Add the signup route
			public.POST("/register", controllers.Register)
		}
		// Add the signup route
		protected := api.Group("/protected").Use(middleware.Auth())
		{
			// Add the profile route
			protected.GET("/profile", controllers.Profile)
		}
	}
	// Return the router
	return r
}
