package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/markysand/ssn"
	"github.com/rickb777/date"
)

func getBirthDateRange(years, months int, now time.Time) (from, to date.Date) {
	switch {
	case years == -1 && months == -1:
		to = date.NewAt(now)
		from = to.AddDate(-100, 0, 0)
	default:
		var (
			intervalYears  = 0
			intervalMonths = 1
			y              = years
			m              = months
		)

		if years == -1 {
			y = 0
		}

		if months == -1 {
			m = 0
			intervalMonths = 0
			intervalYears = 1
		}

		to = date.NewAt(now).AddDate(-y, -m, 0)
		from = to.AddDate(-intervalYears, -intervalMonths, 0)
	}

	return from, to
}

func getRandomBirthDate(from, to date.Date) date.Date {
	diff := to.Sub(from)

	return from.Add(date.PeriodOfDays(rand.IntN(int(diff))) + 1)
}

func main() {
	// parse command line arguments
	flagYears := flag.Int("years", -1, "Set year")
	flagMonths := flag.Int("months", -1, "Set months")
	flagMale := flag.Bool("male", false, "Generate male only")
	flagFemale := flag.Bool("female", false, "Generate female only")
	flagN := flag.Int("n", 1, "Number of ssn-s to generate")
	flag.Usage = func() {
		fmt.Printf(`SSN-Generator is a tool for generating random, safe Swedish SSNs (Social Security Numbers) for testing purposes.

- If you specify months, the generated birth date will have month-level precision.
- If months are not specified, the birth date will have year-level precision.
- If neither years nor months are provided, the generated age will range between 0 and 100 years.

`)
		flag.CommandLine.PrintDefaults()
	}
	flag.Parse()

	// common ssn setup
	lastDigitsArg := []byte("ss?c")

	if *flagFemale {
		lastDigitsArg[2] = 'f'
	}
	if *flagMale {
		lastDigitsArg[2] = 'm'
	}

	from, to := getBirthDateRange(*flagYears, *flagMonths, time.Now())

	// generate ssn-s
	for range *flagN {
		var s ssn.SSN

		birthDate := getRandomBirthDate(from, to)

		s.SetDate(birthDate.In(time.Local))

		s.SetLastDigits(string(lastDigitsArg))

		fmt.Println(s.Format(true, true))
	}
}
