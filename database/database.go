package database

import (
	"github.com/jinzhu/gorm"
)

// Option is an option which is applied to the connection
type Option func(*Connection) error

// WithDebug enables debugging output
func WithDebug() Option {
	return func(c *Connection) error {
		c.db.LogMode(true)
		return nil
	}
}

// Connection is the connection to the database
type Connection struct {
	db *gorm.DB
}

// Connect connects to the database
func Connect(driver, dsn string, options ...Option) (*Connection, error) {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	c := &Connection{
		db: db.Set("gorm:auto_preload", true),
	}

	for _, o := range options {
		o(c)
	}

	c.autoMigrate()
	return c, nil
}

func (c *Connection) autoMigrate() {
	c.db.AutoMigrate(&Org{})
	c.db.AutoMigrate(&ASN{})
	c.db.AutoMigrate(&User{})
	c.db.AutoMigrate(&Tunnel{})
	c.db.AutoMigrate(&Log{})
}

// Close closes the connection
func (c *Connection) Close() error {
	return c.db.Close()
}

// ASNs returns access to the ASN entity
func (c *Connection) ASNs() *ASNDataAccess {
	return &ASNDataAccess{c}
}

// Logs returns access to the Logs entity
func (c *Connection) Logs() *LogDataAccess {
	return &LogDataAccess{c}
}

// Orgs returns access to the Org entity
func (c *Connection) Orgs() *OrgDataAccess {
	return &OrgDataAccess{c}
}

// Tunnels returns access to the Tunnel entity
func (c *Connection) Tunnels() *TunnelDataAccess {
	return &TunnelDataAccess{c}
}

// Users returns access to the User entity
func (c *Connection) Users() *UserDataAccess {
	return &UserDataAccess{c}
}
