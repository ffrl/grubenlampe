package database

import "github.com/jinzhu/gorm"

// Log is the log of all messages processed
type Log struct {
	gorm.Model

	// User is the user requested the operation
	User *User `gorm:"foreignkey:UserID"`

	// UserID is the foreign key to an user
	UserID uint `gorm:"nullable"`

	// Request is the serialized string representation of the request
	RequestMessage string

	// Response is the serialized string representation of the response
	ResponseMessage string

	// Error is the error message
	Error string
}
