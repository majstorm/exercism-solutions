package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	k, _ := time.Parse("1/2/2006 15:04:05", date)
	return k
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	k, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(k)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	k, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour, _, _ := k.Clock()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	formatted_date := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", formatted_date.Weekday(), formatted_date.Month(), formatted_date.Day(), formatted_date.Year(), formatted_date.Hour(), formatted_date.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return Schedule("9/15/2022 00:00:00")
}
