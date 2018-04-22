package database

import "github.com/jinzhu/gorm"

// ASN is an autonomous system number
type ASN struct {
	gorm.Model

	// ASN is the number identifying the autonomous system
	ASN uint32

	// Org is the oranisation holding the AS
	Org *Org `gorm:"foreignkey:OrgID"`

	// OrgID is the foreign key for the organisation
	OrgID uint

	// Checked is the status of the ASN add request
	Checked bool

	// CheckedBy is the authoritative User checked the add request
	CheckedBy *User `gorm:"foreignkey:CheckedByUserID"`

	// CheckedByUserID is the foreign key to the user checked the record
	CheckedByUserID uint `gorm:"nullable"`
}
