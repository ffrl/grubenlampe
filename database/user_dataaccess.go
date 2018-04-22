package database

import "fmt"

// UserDataAccess provides methods to retrieve and store users
type UserDataAccess struct {
	conn *Connection
}

// GetByCredentials gets a user object by a users credentials
func (d *UserDataAccess) GetByCredentials(username, password string) (user *User, err error) {
	err = d.conn.db.Where("email = ? AND password = ?", username, password).First(user).Error
	if err != nil {
		return nil, fmt.Errorf("Query failed: %v", err)
	}

	return user, nil
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
