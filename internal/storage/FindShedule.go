package storage

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/uptrace/bun"

	"sheduler/internal/errorlist"
	"sheduler/internal/models"
)

func (db DB) FindShedule(user_id string, shedule_id string) (models.SheduleWithChart, error) {
	user, err := strconv.Atoi(user_id)
	if err != nil {
		log.Print("Error when converting user_id")
		return models.SheduleWithChart{}, errors.New("Request execution error")
	}
	var shedules []ShedulesStruct
	err = db.conn.NewSelect().Model(&shedules).Where("user_id = ?", user_id).Where("shedule_id = ?", shedule_id).WhereGroup("AND", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.
			Where("status = ?", "actual").
			WhereOr("status = ?", "permanent")
	}).Scan(context.Background())
	if err != nil {
		log.Print(errorlist.ErrorFindShedule)
		return models.SheduleWithChart{}, errors.New("Request execution error")
	}
	if len(shedules) == 0 {
		log.Print(errorlist.ErrorShedulesNotFound)
		return models.SheduleWithChart{}, errors.New(errorlist.ErrorShedulesNotFound)
	}
	for _, value := range shedules {
		var chartUnmarshal models.JSONB
		err := json.Unmarshal(value.ReceptionTiming, &chartUnmarshal)
		if err != nil {
			log.Print(errorlist.ErrorDeserializinJson)
			return models.SheduleWithChart{}, errors.New(errorlist.ErrorDeserializinJson)
		}
		resultModel := models.SheduleWithChart{UserId: int64(user), Medicine: value.MedicineName, Period: value.Period, Duration: value.Duration, Chart: chartUnmarshal}
		return resultModel, nil
	}
	return models.SheduleWithChart{}, errors.New(errorlist.ErrorSearchingShedule)
}
