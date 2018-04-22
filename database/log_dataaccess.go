package database

// LogDataAccess provides methods to retrieve and store logs
type LogDataAccess struct {
	conn *Connection
}

// Insert inserts a log entry
func (d *LogDataAccess) Insert(log *Log) error {
	err := d.conn.db.Create(log).Error
	if err != nil {
		return err
	}

	return nil
}
