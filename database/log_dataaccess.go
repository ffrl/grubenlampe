package database

// LogDataAccess provides methods to retrieve and store logs
type LogDataAccess struct {
	conn *Connection
}

// Insert inserts a log entry
func (d *LogDataAccess) Insert(log *Log) error {
	return d.conn.db.Create(log).Error
}

// GetLog retrieves the full log
func (d *LogDataAccess) GetLog() (ret []*Log, err error) {
	err = d.conn.db.Find(ret).Error
	return ret, err
}
