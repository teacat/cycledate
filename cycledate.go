package cycledate

import (
	"time"
)

type Cycle int

const (
	CycleUnknown Cycle = iota
	CycleDaily
	CycleWeekly
	CycleMonthly
	CycleYearly
)

// lastDay
func lastDay(year int, month time.Month) time.Time {
	firstday := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	return lastday
}

// Next returns the next cycle day.
func Next(t time.Time, cycle Cycle) time.Time {
	//
	year, month, _ := t.Date()
	hour, minute, second := t.Clock()
	//
	switch cycle {
	case CycleDaily:
		return t.Add(time.Hour * 24)
	case CycleWeekly:
		return t.Add(time.Hour * 24 * 7)
	case CycleMonthly:
		nextCycleLastDay := lastDay(year, month+1)
		// If the day of today doesn't exist in the next month (e.g. 01/31 > 02/28)
		if t.Day() > nextCycleLastDay.Day() {
			return time.Date(year, month+1, nextCycleLastDay.Day(), hour, minute, second, 0, t.Location())
		} else {
			return time.Date(year, month+1, t.Day(), hour, minute, second, 0, t.Location())
		}
	case CycleYearly:
		nextCycleLastDay := lastDay(year+1, month)
		// If the day of today doesn't exist in the next year (e.g. 01/31 > 02/28)
		if t.Day() > nextCycleLastDay.Day() {
			return time.Date(year+1, month, nextCycleLastDay.Day(), hour, minute, second, 0, t.Location())
		} else {
			return time.Date(year+1, month, t.Day(), hour, minute, second, 0, t.Location())
		}
	}
	return t
}

// NextZero calls Next then set the next cycle date happened in 00:00 AM instead of the same time(hour/minute) as last cycle.
// Useful if you want to group up all the cycle that will happened in the same day.
func NextZero(t time.Time, cycle Cycle) time.Time {
	year, month, day := Next(t, cycle).Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
