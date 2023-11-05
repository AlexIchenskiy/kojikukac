package models

import (
	"backend/database"

	"gorm.io/gorm"
)

// in order to keep track of sessions a user is logged into, we need a table to keep track of their emails and tokens, furthermore, since a user might have multiple sessions, the emails
// should not be unique
type Session struct {
	gorm.Model
	Email string `json:"email" binding:"required"`
    Token string `json:"token" binding:"required" gorm:"unique"`
}

func (session *Session) CreateSessionRecord() error {
	result := database.GlobalDB.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
