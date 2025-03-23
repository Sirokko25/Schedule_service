package storage

import (
	"context"
	"log"
	"time"
)

func (db DB) Actualize() {
	var shedules []ShedulesStruct
	now := time.Now().UTC()
	err := db.conn.NewSelect().Model(&shedules).Where("status = ?", "actual").Where("end_date < ?", now).Scan(context.Background())
	if err != nil {
		log.Printf("Error checking the relevance of schedules: %v", err)
		return
	}
	if len(shedules) == 0 {
		log.Print("There are no outdated schedules")
		return
	}
	for _, shedule := range shedules {
		storageObject := &HistoryShedulesStruct{SheduleId: shedule.SheduleId, MedicineName: shedule.MedicineName, UserId: shedule.UserId, StartDate: shedule.StartDate, EndDate: shedule.EndDate}
		var id int64
		_, err = db.conn.NewInsert().Model(storageObject).Returning("shedule_id").Exec(context.Background(), &id)
		if err != nil {
			log.Printf("Error when moving an entry %v", err)
			return
		}
		ctx := context.Background()
		result, err := db.conn.NewDelete().Model(&shedules).Where("shedule_id = ?", shedule.SheduleId).Exec(ctx)
		if err != nil {
			log.Printf("Error deleting the schedule: %v", err)
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			log.Print("No records were found.")
		}
		log.Printf("Rows deleted: %d", rowsAffected)
	}
	return
}
