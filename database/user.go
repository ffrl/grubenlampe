package database

// User is an user account permited to use the API
type User struct {
	// ID identifies an User
	ID int

	// Email identifies the user in API communications
	Email string

	// Password is password of the User
	Password string

	// SuperUser is the role of an User permited to perform admin actions
	SuperUser bool

	// RIPEHandle is the reference in the RIPE database
	RIPEHandle string
}
