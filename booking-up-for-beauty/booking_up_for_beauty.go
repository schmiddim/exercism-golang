package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.

func Schedule(date string) time.Time {
	//Mon Jan 2 15:04:05 -0700 MST 2006.
	const layout = "1/2/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	//Mon Jan 2 15:04:05 -0700 MST 2006.
	const layout = "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return t.Before(time.Now())

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	//Mon Jan 2 15:04:05 -0700 MST 2006.
	const layout = "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}

	afterNoonStart := time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, t.Location())
	afterNoonEnd := time.Date(t.Year(), t.Month(), t.Day(), 18, 0, 0, 0, t.Location())

	return t.After(afterNoonStart) && t.Before(afterNoonEnd)

}

// Description returns a formatted string of the appointment time.
func Description(date string) string {

	t := Schedule(date)

	msg := "You have an appointment on %s, %s %d, %d, at %d:%d."
	return fmt.Sprintf(msg, t.Weekday(), t.Month().String(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() (t time.Time) {

	now := time.Now()
	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, t.Location())
}
