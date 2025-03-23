package storage

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"sheduler/internal/errorlist"
	"sheduler/internal/helpers"
	"sheduler/internal/models"
)

func (db DB) AppendShedule(Shedule models.Shedule) (int64, error) {
	now := time.Now().UTC()
	if Shedule.Duration == "permanent" {
		resultInterval := helpers.CreateIntervals(now, Shedule.Period)
		resultIntervalsSlice, err := json.Marshal(resultInterval)
		if err != nil {
			log.Print(errorlist.ErrorSerializinJson)
			return 0, errors.New(errorlist.ErrorSerializinJson)
		}
		storageObject := &ShedulesStruct{MedicineName: Shedule.Medicine, UserId: Shedule.UserId, Period: Shedule.Period, Duration: Shedule.Duration, EndDate: time.Time{}, Status: "permanent", ReceptionTiming: resultIntervalsSlice}
		var id int64
		_, err = db.conn.NewInsert().Model(storageObject).Returning("shedule_id").Exec(context.Background(), &id)
		if err != nil {
			log.Printf("Error when adding a record to the database %v", err)
			return 0, errors.New("Error when adding a record to the database")
		}
		return id, nil
	} else {
		resultDate, err := helpers.EndDateCalculate(now, Shedule.Duration)
		if err != nil {
			log.Print("Error in calculating the end date")
			return 0, errors.New("Error in calculating the end date")
		}
		resultInterval := helpers.CreateIntervals(now, Shedule.Period)
		resultIntervalsSlice, err := json.Marshal(resultInterval)
		if err != nil {
			log.Print("errorlist.ErrorSerializinJson")
			return 0, errors.New("errorlist.ErrorSerializinJson")
		}
		storageObject := &ShedulesStruct{MedicineName: Shedule.Medicine, UserId: Shedule.UserId, Period: Shedule.Period, Duration: Shedule.Duration, EndDate: resultDate, Status: "actual", ReceptionTiming: resultIntervalsSlice}
		var id int64
		_, err = db.conn.NewInsert().Model(storageObject).Returning("shedule_id").Exec(context.Background(), &id)
		if err != nil {
			log.Print("Error adding an entry")
			return 0, errors.New("Error adding an entry")
		}
		return id, nil
	}
}
