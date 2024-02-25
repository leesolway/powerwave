package utils

import "time"

// DaysInMonth calculates the number of days in the month for a given date.
func DaysInMonth(date time.Time) int {
	firstOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1) // Move to the first day of the next month, then subtract one day
	return lastOfMonth.Day()
}
