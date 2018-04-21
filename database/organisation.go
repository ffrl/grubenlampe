package database

// Org is an organisation
type Org struct {
	// ID identifies an organisation
	ID int

	// Address is the address of the organisation
	Address string

	// Short name is the name used to identify the organisation in backend systems
	ShortName string

	// Active is the current status of the organisation
	Active bool

	// Checked is the status of the organisation add request
	Checked bool

	// CheckedBy is the authorative user checked the add request
	CheckedBy *User

	// IPv4Quota is the max number of NAT-IPs an organisation can hold
	IPv4Quota uint8

	// IPv6Quota is the max number of /48 prefixes an organisation can hold
	IPv6Quota uint8
}
