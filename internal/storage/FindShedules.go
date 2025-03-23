package storage

import (
	"context"
	"errors"
	"log"

	"github.com/uptrace/bun"

	"sheduler/internal/errorlist"
)

func (db DB) FindShedules(user_id string) ([]int64, error) {
	var shedulesID []int64
	var shedules []ShedulesStruct
	err := db.conn.NewSelect().Model(&shedules).Where("user_id = ?", user_id).WhereGroup("AND", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.
			Where("status = ?", "actual").
			WhereOr("status = ?", "permanent")
	}).Scan(context.Background())
	if err != nil {
		log.Print("Error when searching for a schedule")
		return nil, errors.New("Error when searching for a schedule")
	}
	if len(shedules) == 0 {
		log.Print(errorlist.ErrorShedulesNotFound)
		return nil, errors.New(errorlist.ErrorShedulesNotFound)
	}
	for _, shedule := range shedules {
		shedulesID = append(shedulesID, shedule.SheduleId)
	}
	log.Print("Schedules have been successfully returned")
	return shedulesID, nil
}
