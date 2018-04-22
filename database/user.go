package database

import "github.com/jinzhu/gorm"

// User is an user account permited to use the API
type User struct {
	gorm.Model

	// Email identifies the user in API communications
	Email string

	// Password is password of the user
	Password string

	// SuperUser is the role of an user permited to perform admin actions
	SuperUser bool

	// RIPEHandle is the reference in RIPE database
	RIPEHandle string `gorm:"column:ripe_handle"`

	// Orgs is the list of organisations an user is assinged to
	Orgs []*Org `gorm:"many2many:user_org"`
}

// HasOrg checks if a user belongs to ordID Org
func (u *User) HasOrg(orgID uint) bool {
	for _, org := range u.Orgs {
		if org.ID == orgID {
			return true
		}
	}
	return false
}
