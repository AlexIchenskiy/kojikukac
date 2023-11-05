package controllers

import (
	"backend/auth"
	"backend/config"
	"backend/database"
	"backend/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPayload login body
// LoginPayload is a struct that contains the fields for a user's login credentials
type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse token response
// LoginResponse is a struct that contains the fields for a user's login response
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

// Login is a function that handles user login
// It takes in a gin context as an argument and binds the user data from the request body to a LoginPayload struct
// It then checks if the user exists in the database and if the password is correct
// If successful, it generates a token and a refresh token and returns a 200 status code with the token and refresh token
// If unsuccessful, it returns a 401 or 500 status code with an error message

func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	for _, admin := range config.GetAdmins() {
		if admin.Email == payload.Email && admin.Password == payload.Password {
			jwtWrapper := auth.JwtWrapper{
				SecretKey:         "verysecretkey",
				Issuer:            "AuthService",
				ExpirationMinutes: 1,
				ExpirationHours:   12,
			}
			signedToken, err := jwtWrapper.GenerateToken(user.Email)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"Error": "Error Signing Token",
				})
				c.Abort()
				return
			}
			signedtoken, err := jwtWrapper.RefreshToken(user.Email)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"Error": "Error Signing Token",
				})
				c.Abort()
				return
			}
			tokenResponse := LoginResponse{
				Token:        signedToken,
				RefreshToken: signedtoken,
			}
			config.AdminTokens = append(config.AdminTokens, tokenResponse.Token)
			c.JSON(200, tokenResponse)
		}
	}

	result := database.GlobalDB.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}
	signedToken, err := jwtWrapper.GenerateToken(fmt.Sprintf("%s%s", user.Email, time.Now()))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}

    // create a new session record
    newSession := models.Session {
        Email: user.Email,
        Token: signedToken,
    }
    newSession.CreateSessionRecord()

	tokenResponse := LoginResponse{
		Token: signedToken,
	}
	c.JSON(200, tokenResponse)
}
