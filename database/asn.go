package database

// ASN is an autonomous system number
type ASN struct {
	// ID identifies an autonomous system
	ID int

	// ASN is the number identifying the autonomous system
	ASN uint32

	// Org is the oranisation holding the AS
	Org *Org

	// Checked is the status of the ASN add request
	Checked bool

	// CheckedBy is the authorative User checked the add request
	CheckedBy *User
}
