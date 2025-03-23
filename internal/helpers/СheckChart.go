package helpers

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"sheduler/internal/models"
)

func CheckChart(chart models.JSONB) ([]string, error) {
	resultSlice := make([]string, 0)
	now := time.Now().UTC()
	for _, v := range chart.Time {
		splitChart := strings.Split(v, ":")
		hours, err := strconv.Atoi(splitChart[0])
		if err != nil {
			log.Printf("Error when converting hours %s: %v", v, err)
			return nil, errors.New("Error when converting hours")
		}
		minutes, err := strconv.Atoi(splitChart[1])
		if err != nil {
			log.Printf("Error when converting minutes %s: %v", v, err)
			return nil, errors.New("Error when converting minutes")
		}
		IntervalTime := time.Date(now.Year(), now.Month(), now.Day(), hours, minutes, 0, 0, time.UTC)
		timeDifference := IntervalTime.Sub(now)
		if timeDifference <= time.Hour && timeDifference > 0 {
			resultSlice = append(resultSlice, v)
		} else {
			continue
		}
	}
	return resultSlice, nil
}
