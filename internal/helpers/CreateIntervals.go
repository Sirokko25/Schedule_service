package helpers

import (
	"time"

	"sheduler/internal/models"
)

func CreateIntervals(now time.Time, minutes int64) models.JSONB {
	if minutes >= 1440 && minutes%15 == 0 {
		return CalculateDayIntervals(minutes, now)

	} else if minutes >= 1440 {
		minutes = int64((minutes/15)+1) * 15
		now = now.Add(time.Duration(minutes) * time.Minute)
		return CalculateDayIntervals(minutes, now)

	} else if minutes%15 == 0 && float64(minutes)/15 > 1 {
		minutes += 0

	} else if float64(minutes)/15 > 1 {
		minutes = int64((minutes/15)+1) * 15

	} else {
		minutes = 15
	}
	return CalculateMinutesIntervals(minutes, now)
}
