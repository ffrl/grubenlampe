package database

import "github.com/jinzhu/gorm"

// Tunnel is a GRE tunnel connecting an organisation to backbone routers
type Tunnel struct {
	gorm.Model

	// ASN identifies the ASN requesting the tunnel
	ASN *ASN `gorm:"foreignkey:ASNID"`

	// ASNID is the foreign key to an ASN
	ASNID uint

	// Address is the IP of the GRE endpoint (organisation site)
	Address string

	// Synced describes if a tunnel is synced to Netbox for provisioning
	Synced bool
}
