package utils

import "time"

func GetTotalHoursInAMonth(year int, month int) int {
	// Get the current time
	currentTime := time.Now()
	// Get the first day of the current month
	firstDayOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, currentTime.Location())
	// Calculate the last day of the current month
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	// Calculate the total number of days in the current month
	return lastDayOfMonth.Day() * 24

}
