package main

import (
	"time"

	"github.com/rickb777/date"
)

func getBirthRangeFromRange(now date.Date, years, months int) (from, to date.Date) {
	var (
		intervalYears      = 0
		intervalMonths     = 1
		y              int = years
		m              int = months
	)

	if years == -1 {
		y = 0
	}

	if months == -1 {
		m = 0
		intervalMonths = 0
		intervalYears = 1
	}

	to = now.AddDate(-y, -m, 0)
	from = to.AddDate(-intervalYears, -intervalMonths, 0)

	return from, to
}

func getBirthRange(c *config) (from, to date.Date) {
	now := date.NewAt(time.Now())

	switch {
	case c.child:
		to = now
		from = to.AddDate(-teenAge, 0, 0)

	case c.child:
		to = now.AddDate(-teenAge, 0, 0)
		from = now.AddDate(-adultAge, 0, 0)

	case c.adult:
		to = now.AddDate(-adultAge, 0, 0)
		from = now.AddDate(-maxAge, 0, 0)

	case c.years == -1 && c.months == -1:
		to = now
		from = to.AddDate(-maxAge, 0, 0)

	default:
		from, to = getBirthRangeFromRange(now, c.years, c.months)
	}

	return from, to
}
