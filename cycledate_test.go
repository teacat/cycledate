package cycledate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// EqualTime
func EqualTime(as *assert.Assertions, expected, actual time.Time) {
	y, m, d := actual.Date()
	h, i, s := actual.Clock()
	as.Equal(expected.Year(), y)
	as.Equal(expected.Month(), m)
	as.Equal(expected.Day(), d)
	as.Equal(expected.Hour(), h)
	as.Equal(expected.Minute(), i)
	as.Equal(expected.Second(), s)
}

func TestDaily(test *testing.T) {
	a := assert.New(test)
	// 2024-01-30 -> 2024-01-31
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-01-30 06:30:30")
	EqualTime(a, time.Date(2024, time.January, 31, 6, 30, 30, 0, t.Location()), Next(t, CycleDaily))
	// 2024-01-31 -> 2024-02-01
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-01-31 06:30:30")
	EqualTime(a, time.Date(2024, time.February, 1, 6, 30, 30, 0, t.Location()), Next(t, CycleDaily))
}

func TestWeekly(test *testing.T) {
	a := assert.New(test)
	// 2024-01-31 -> 2024-02-07
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-01-31 06:30:30")
	EqualTime(a, time.Date(2024, time.February, 7, 6, 30, 30, 0, t.Location()), Next(t, CycleWeekly))
}

func TestMonthly(test *testing.T) {
	a := assert.New(test)
	// 2024-01-31 -> 2024-02-29; 31 -> 29
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-01-31 06:30:30")
	EqualTime(a, time.Date(2024, time.February, 29, 6, 30, 30, 0, t.Location()), Next(t, CycleMonthly))
	// 2024-02-29 -> 2024-03-29; 29 -> 29
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-02-29 06:30:30")
	EqualTime(a, time.Date(2024, time.March, 29, 6, 30, 30, 0, t.Location()), Next(t, CycleMonthly))
	// 2024-03-29 -> 2024-04-29; 29 -> 29
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-03-29 06:30:30")
	EqualTime(a, time.Date(2024, time.April, 29, 6, 30, 30, 0, t.Location()), Next(t, CycleMonthly))
}

func TestYearly(test *testing.T) {
	a := assert.New(test)
	// 2024-02-29 -> 2025-02-28; Leap Year -> Normal Year
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-02-29 06:30:30")
	EqualTime(a, time.Date(2025, time.February, 28, 6, 30, 30, 0, t.Location()), Next(t, CycleYearly))
	// 2024-01-29 -> 2025-01-29; 29 -> 29
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-01-29 06:30:30")
	EqualTime(a, time.Date(2025, time.January, 29, 6, 30, 30, 0, t.Location()), Next(t, CycleYearly))
}

func TestDailyZero(test *testing.T) {
	a := assert.New(test)
	// 2024-01-30 -> 2024-01-31
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-01-30 06:30:30")
	EqualTime(a, time.Date(2024, time.January, 31, 0, 0, 0, 0, t.Location()), NextZero(t, CycleDaily))
	// 2024-01-31 -> 2024-02-01
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-01-31 06:30:30")
	EqualTime(a, time.Date(2024, time.February, 1, 0, 0, 0, 0, t.Location()), NextZero(t, CycleDaily))
}

func TestWeeklyZero(test *testing.T) {
	a := assert.New(test)
	// 2024-01-31 -> 2024-02-07
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-01-31 06:30:30")
	EqualTime(a, time.Date(2024, time.February, 7, 0, 0, 0, 0, t.Location()), NextZero(t, CycleWeekly))
}

func TestMonthlyZero(test *testing.T) {
	a := assert.New(test)
	// 2024-01-31 -> 2024-02-29; 31 -> 29
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-01-31 06:30:30")
	EqualTime(a, time.Date(2024, time.February, 29, 0, 0, 0, 0, t.Location()), NextZero(t, CycleMonthly))
	// 2024-02-29 -> 2024-03-29; 29 -> 29
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-02-29 06:30:30")
	EqualTime(a, time.Date(2024, time.March, 29, 0, 0, 0, 0, t.Location()), NextZero(t, CycleMonthly))
	// 2024-03-29 -> 2024-04-29; 29 -> 29
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-03-29 06:30:30")
	EqualTime(a, time.Date(2024, time.April, 29, 0, 0, 0, 0, t.Location()), NextZero(t, CycleMonthly))
}

func TestYearlyZero(test *testing.T) {
	a := assert.New(test)
	// 2024-02-29 -> 2025-02-28; Leap Year -> Normal Year
	t, _ := time.Parse("2006-01-02 03:04:05", "2024-02-29 06:30:30")
	EqualTime(a, time.Date(2025, time.February, 28, 0, 0, 0, 0, t.Location()), NextZero(t, CycleYearly))
	// 2024-01-29 -> 2025-01-29; 29 -> 29
	t, _ = time.Parse("2006-01-02 03:04:05", "2024-01-29 06:30:30")
	EqualTime(a, time.Date(2025, time.January, 29, 0, 0, 0, 0, t.Location()), NextZero(t, CycleYearly))
}
