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

func (reservation *Resevation) CreateReservation() error {
	result := database.GlobalDB.Create(&reservation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
