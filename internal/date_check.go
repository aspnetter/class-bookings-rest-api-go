package internal

import "time"

func DateInRange(date time.Time, startDate time.Time, endDate time.Time) bool {
	return date.After(startDate.AddDate(0, 0, -1)) && date.Before(endDate.AddDate(0, 0, 1))
}
