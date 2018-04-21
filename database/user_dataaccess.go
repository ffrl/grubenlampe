package database

// UserDataAccess provides methods to retrieve and store users
type UserDataAccess struct {
	conn *Connection
}

// Verify verfies user credentials
func (d *UserDataAccess) Verify(username, password string) (bool, error) {
	var count int
	err := d.conn.db.Model(&User{}).
		Where("email = ? AND password = ?", username, password).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
