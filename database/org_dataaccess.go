package database

type OrgDataAccess struct {
	conn *Connection
}

// Save persists an organisation
func (d *OrgDataAccess) Save(o *Org) error {
	return d.conn.db.Save(o).Error
}

// GetByShortName retrieves an organisation by its short name
func (d *OrgDataAccess) GetByShortName(name string) (*Org, error) {
	o := &Org{}
	err := d.conn.db.First(&o, "short_name = ?", name).Error

	return o, err
}

// ShortNameExists checks if an short name is already taken
func (d *OrgDataAccess) ShortNameExists(name string) (bool, error) {
	var count int
	err := d.conn.db.Model(&Org{}).
		Where("short_name = ?", name).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
