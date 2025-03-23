package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"sheduler/internal/errorlist"
)

func ConnectionDB() (DB, error) {

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DSN"))))
	if sqldb == nil {
		log.Fatal("Failed to create connection")
		return DB{}, errors.New(errorlist.ErrorCreateTables)
	}
	bunConn := bun.NewDB(sqldb, pgdialect.New())
	ctx := context.Background()
	_, err := bunConn.NewCreateTable().Model((*ShedulesStruct)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatalf("Failed to create table shedules: %v", err)
		return DB{}, errors.New(errorlist.ErrorCreateTables)
	}
	log.Print("The ShedulesStruct table has been created")
	_, err = bunConn.NewCreateTable().Model((*HistoryShedulesStruct)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatalf("Failed to create table shedules_history: %v", err)
		return DB{}, errors.New(errorlist.ErrorCreateTables)
	}
	log.Print("The HistoryShedulesStruct table has been created")
	log.Print("The connection has been created")
	return DB{conn: bunConn}, nil
}
