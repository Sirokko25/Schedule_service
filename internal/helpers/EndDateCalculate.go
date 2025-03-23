package helpers

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func EndDateCalculate(now time.Time, durationString string) (time.Time, error) {
	durationSlice := strings.Split(durationString, "")
	switch durationSlice[1] {
	case "d":
		days, err := strconv.Atoi(durationSlice[0])
		if err != nil {
			return time.Time{}, errors.New("Error in calculating the end date of medication")
		}
		return now.AddDate(0, 0, days), nil
	case "w":
		days, err := strconv.Atoi(durationSlice[0])
		if err != nil {
			return time.Time{}, errors.New("Error in calculating the end date of medication")
		}
		return now.AddDate(0, 0, days*7), nil
	case "m":
		months, err := strconv.Atoi(durationSlice[0])
		if err != nil {
			return time.Time{}, errors.New("Error in calculating the end date of medication")
		}
		return now.AddDate(0, months, 0), nil
	case "y":
		years, err := strconv.Atoi(durationSlice[0])
		if err != nil {
			return time.Time{}, errors.New("Error in calculating the end date of medication")
		}
		return now.AddDate(years, 0, 0), nil
	}
	return time.Time{}, errors.New("Error in calculating the end date of medication")
}
