package storage

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/uptrace/bun"

	"sheduler/internal/errorlist"
	"sheduler/internal/helpers"
	"sheduler/internal/models"
)

func (db DB) NextTakings(user_id string) ([]models.Medicine, error) {
	var MedicineResult []models.Medicine
	var shedules []ShedulesStruct
	err := db.conn.NewSelect().Model(&shedules).Where("user_id = ?", user_id).WhereGroup("AND", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.
			Where("status = ?", "actual").
			WhereOr("status = ?", "permanent")
	}).Scan(context.Background())
	if err != nil {
		log.Printf("%s: %v", errorlist.ErrorSearchingShedule, err)
		return nil, errors.New(errorlist.ErrorSearchingShedule)
	}
	if len(shedules) == 0 {
		log.Print(errorlist.ErrorShedulesNotFound)
		return nil, errors.New(errorlist.ErrorShedulesNotFound)
	}
	for _, shedule := range shedules {
		var chartUnmarshal models.JSONB
		err := json.Unmarshal(shedule.ReceptionTiming, &chartUnmarshal)
		if err != nil {
			log.Print(errorlist.ErrorDeserializinJson)
			return nil, errors.New(errorlist.ErrorDeserializinJson)
		}
		time, err := helpers.CheckChart(chartUnmarshal)
		if err != nil {
			log.Print("Error when checking the nearest medication intake time")
			return nil, errors.New("Error when checking the nearest medication intake time")
		}
		if len(time) == 0 {
			continue
		} else {
			Medicine := models.Medicine{Medicine: shedule.MedicineName, Time: time}
			MedicineResult = append(MedicineResult, Medicine)
		}
	}
	if len(MedicineResult) == 0 {
		log.Print("There are no pills needed to take in the near future")
		return nil, errors.New("There are no pills needed to take in the near future")
	}
	return MedicineResult, nil
}
