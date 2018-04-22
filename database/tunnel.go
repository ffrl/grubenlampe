package database

import "github.com/jinzhu/gorm"

// Tunnel is a GRE tunnel connecting an organisation to backbone routers
type Tunnel struct {
	gorm.Model

	// ASN identifies the ASN requesting the tunnel
	ASN *ASN `gorm:"foreignkey:ASNID"`

	// ASNID is the foreign key to an ASN
	ASNID uint

	// Router is the identification string for a router
	Router string

	// Address is the IP of the GRE endpoint (organisation site)
	Address string

	// Synced describes if a tunnel is synced to Netbox for provisioning
	Synced bool
}
