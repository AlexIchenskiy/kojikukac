package models

import (
	"time"
)

type ParkingSpot struct {
	ID                string    `json:"id"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	ParkingSpotZone   string    `json:"parkingSpotZone"`
	Occupied          bool      `json:"occupied"`
	OccupiedTimestamp time.Time `json:"occupiedTimestamp"`
}

var ParkingSpots []ParkingSpot
