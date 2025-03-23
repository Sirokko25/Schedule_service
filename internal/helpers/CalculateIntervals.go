package helpers

import (
	"time"

	"sheduler/internal/models"
)

func CalculateDayIntervals(minutes int64, now time.Time) models.JSONB {
	now = now.Add(time.Duration(minutes+15) * time.Minute)
	if now.Hour() > 22 || now.Hour() < 8 {
		now = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.UTC)
	}
	formattedTime := now.Format("15:04")
	resultSlice := models.JSONB{Time: []string{formattedTime}}
	return resultSlice
}

func CalculateMinutesIntervals(minutes int64, now time.Time) models.JSONB {
	resultTimeSlice := make([]string, 0)
	timeStart := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.UTC)
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 1, 0, time.UTC)
	for timeStart.Before(endTime) {
		formattedTime := timeStart.Format("15:04")
		resultTimeSlice = append(resultTimeSlice, formattedTime)

		timeStart = timeStart.Add(time.Duration(minutes) * time.Minute)
	}
	return models.JSONB{Time: resultTimeSlice}

}
