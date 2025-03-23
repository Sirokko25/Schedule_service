package storage

import "context"

func (db DB) PingDB(context context.Context) error {
	err := db.conn.DB.PingContext(context)
	return err
}
