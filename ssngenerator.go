package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/markysand/ssn"
	"github.com/rickb777/date"
)

const (
	teenAge  = 13
	adultAge = 18
	maxAge   = 100
)

func getRandomBirthDate(from, to date.Date) date.Date {
	diff := to.Sub(from)

	return from.Add(date.PeriodOfDays(rand.IntN(int(diff))) + 1)
}

func main() {
	flagYears := flag.Int("years", -1, "Set year")
	flagMonths := flag.Int("months", -1, "Set months")
	flagChild := flag.Bool("child", false, "Child, 0-12 years")
	flagTeen := flag.Bool("teen", false, "Teen, 13-17 years")
	flagAdult := flag.Bool("adult", false, "Adult, 18-100 years")
	flagMale := flag.Bool("male", false, "Generate male only")
	flagFemale := flag.Bool("female", false, "Generate female only")
	flagN := flag.Int("n", 1, "Number of ssn-s to generate")

	flag.Usage = func() {
		fmt.Printf(`SSN-Generator is a tool for generating random, safe Swedish SSNs (Social Security Numbers) for testing purposes.

- If you specify months, the generated birth date will have month-level precision.
- If months are not specified, the birth date will have year-level precision.
- If neither years nor months are provided, the generated age will range between 0 and 100 years.
- Child, teen, adult will generate ages in the ranges of 0-12, 13-17, 18-100 years respectively.

`)
		flag.CommandLine.PrintDefaults()
	}

	flag.Parse()

	lastDigitsArg := []byte("ss?c")

	if *flagFemale {
		lastDigitsArg[2] = 'f'
	}
	if *flagMale {
		lastDigitsArg[2] = 'm'
	}

	var from, to date.Date

	now := date.NewAt(time.Now())

	switch {
	case *flagChild:
		to = now
		from = to.AddDate(-teenAge, 0, 0)

	case *flagTeen:
		to = now.AddDate(-teenAge, 0, 0)
		from = now.AddDate(-adultAge, 0, 0)

	case *flagAdult:
		to = now.AddDate(-adultAge, 0, 0)
		from = now.AddDate(-maxAge, 0, 0)

	case *flagYears == -1 && *flagMonths == -1:
		to = now
		from = to.AddDate(-maxAge, 0, 0)

	default:
		var (
			intervalYears      = 0
			intervalMonths     = 1
			y              int = *flagYears
			m              int = *flagMonths
		)

		if *flagYears == -1 {
			y = 0
		}

		if *flagMonths == -1 {
			m = 0
			intervalMonths = 0
			intervalYears = 1
		}

		to = now.AddDate(-y, -m, 0)
		from = to.AddDate(-intervalYears, -intervalMonths, 0)
	}

	for range *flagN {
		birthDate := getRandomBirthDate(from, to)

		var s ssn.SSN
		s.SetDate(birthDate.In(time.Local))
		s.SetLastDigits(string(lastDigitsArg))

		fmt.Println(s.Format(true, true))

	}
}
