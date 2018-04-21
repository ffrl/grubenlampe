package database

import "github.com/jinzhu/gorm"

// Org is an organisation
type Org struct {
	gorm.Model

	// Name is the long name of the organisation
	Name string

	// Address is the address of the organisation
	Address string

	// Short name is the name used to identify the organisation in backend systems
	ShortName string `gorm:"size:5"`

	// Active is the current status of the organisation
	Active bool

	// Checked is the status of the organisation add request
	Checked bool

	// CheckedBy is the authoritative user checked the add request
	CheckedBy *User

	// IPv4Quota is the max number of NAT-IPs an organisation can hold
	IPv4Quota uint8 `gorm:"column:ipv4_quota"`

	// IPv6Quota is the max number of /48 prefixes an organisation can hold
	IPv6Quota uint8 `gorm:"column:ipv6_quota"`
}
