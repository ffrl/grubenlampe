package database

// User is an user account permited to use the API
type User struct {
	// ID identifies an user
	ID int

	// Email identifies the user in API communications
	Email string

	// Password is password of the user
	Password string

	// SuperUser is the role of an user permited to perform admin actions
	SuperUser bool

	// RIPEHandle is the reference in RIPE database
	RIPEHandle string
}
