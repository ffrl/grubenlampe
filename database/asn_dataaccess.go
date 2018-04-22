package database

// ASNDataAccess provides methods to retrieve and store ASNs
type ASNDataAccess struct {
	conn *Connection
}

// GetByNumber retrieves the ASN for a number
func (d *ASNDataAccess) GetByNumber(asn uint32) (*ASN, error) {
	a := &ASN{}
	err := d.conn.db.Preload("Org").First(&a, "asn = ?", asn).Error
	if err != nil {
		return nil, err
	}

	return a, err
}

// Save persists an ASN
func (d *ASNDataAccess) Save(a *ASN) error {
	return d.conn.db.Save(a).Error
}

// CheckedASNExists checks if an checked ASN record exists for a given AS number
func (d *ASNDataAccess) CheckedASNExists(a uint32) (bool, error) {
	var count int
	err := d.conn.db.Model(&ASN{}).
		Where("asn = ? AND checked", a).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
