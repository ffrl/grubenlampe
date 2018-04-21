package database

import "github.com/jinzhu/gorm"

// Log is the log of all messages processed
type Log struct {
	gorm.Model

	// User is the user requested the operation
	User *User

	// Org is the organisation on which behalv the user requested the action
	Org *Org

	// Request is the serialized string representation of the request
	RequestMessage string

	// Response is the serialized string representation of the response
	ResponseMessage string
}
