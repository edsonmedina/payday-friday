package payday

import (
	"time"
)

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// LastFridayOf returns the day for the last friday of the month
func LastFridayOf(m time.Month, year int) int {
	numDaysInMonth := daysIn(m, year)

	lastDay := time.Date(year, m, numDaysInMonth, 0, 0, 0, 0, time.UTC)
	weekday := int(lastDay.Weekday())

	if weekday == 5 { //"Friday"
		return lastDay.Day()
	}

	if weekday == 6 {
		return numDaysInMonth - 1
	}

	return numDaysInMonth - 7 + 5 - weekday
}

// WeeksLong returns the number of weeks between the payday of a given month and the previous
func WeeksLong(m time.Month, year int) int {
	lastFriday := LastFridayOf(m, year)
	paydayTime := time.Date(year, m, lastFriday, 0, 0, 0, 0, time.UTC)

	previousMonth := m - 1

	previousLastFriday := LastFridayOf(previousMonth, year)
	previousPaydayTime := time.Date(year, previousMonth, previousLastFriday, 0, 0, 0, 0, time.UTC)

	difference := paydayTime.Sub(previousPaydayTime)

	return int(difference.Hours()) / 24 / 7
}
