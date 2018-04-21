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

// Save persists an user
func (d *UserDataAccess) Save(u *User) error {
	return d.conn.db.Save(u).Error
}

// EmailExists checks if an email address already exists
func (d *UserDataAccess) EmailExists(email string) (bool, error) {
	var count int
	err := d.conn.db.Model(&User{}).
		Where("email = ?", email).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetByEmail retrieves an user by its email address
func (d *UserDataAccess) GetByEmail(email string) (*User, error) {
	u := &User{}
	err := d.conn.db.First(&u, "email = ?", email).Error

	return u, err
}
