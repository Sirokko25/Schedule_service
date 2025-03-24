package storage

import (
	"context"
	"log"
)

func (db DB) PingDB(context context.Context) error {
	err := db.conn.DB.PingContext(context)
	log.Print("The database connection is active")
	return err
}
