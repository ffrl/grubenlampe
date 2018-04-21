package database

// Org is an organisation
type Org struct {
	// ID identifies an Org
	ID int

	// Address is the address of the Org
	Address string

	// Short name is the name used to identify the Org in backend systems
	ShortName string

	// Active is the current status of the Org
	Active bool

	// Checked is the status of the Org add request
	Checked bool

	// CheckedBy is the authorative user checked the add request
	CheckedBy *User

	// IPv4Quota is the max number of NAT-IPs an Org can hold
	IPv4Quota uint8

	// IPv6Quota is the max number of /48 prefixes an Org can hold
	IPv6Quota uint8
}
