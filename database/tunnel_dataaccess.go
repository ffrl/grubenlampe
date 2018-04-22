package database

import "fmt"

// TunnelDataAccess provides methods to retrieve and store tunnels
type TunnelDataAccess struct {
	conn *Connection
}

func (t *TunnelDataAccess) GetTunnelsByAddress(address string) (ret []*Tunnel, err error) {
	err = t.conn.db.Where("address = ?", address).Find(ret).Error
	if err != nil {
		return nil, fmt.Errorf("Unable to get data: %v", err)
	}

	return ret, nil
}

func (t *TunnelDataAccess) AddTunnel(asn uint32, address string) error {
	err := t.conn.db.Create(Tunnel{
		ASN:     &ASN{ASN: asn},
		Address: address,
	}).Error
	if err != nil {
		return fmt.Errorf("Unable to write data: %v", err)
	}

	return nil
}
