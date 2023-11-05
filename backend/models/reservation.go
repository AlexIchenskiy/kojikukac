package models

import (
	"backend/database"

	"gorm.io/gorm"
)

type Resevation struct {
	gorm.Model
	Email         string `json:"email"  gorm:"unique"`
	ParkingSpotId string `json:"parkingspotid" binding:"required" gorm:"primaryKey"`
	EndH          string `json:"endh" binding:"required"`
	EndM          string `json:"endm" binding:"required"`
}

// CreateUserRecord creates a user record in the database
// CreateUserRecord takes a pointer to a User struct and creates a user record in the database
// It returns an error if there is an issue creating the user record
func (rezervation *Resevation) CreateReservation() error {
	result := database.GlobalDB.Create(&rezervation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
