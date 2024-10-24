package main

import (
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

func lastDigits(c *config) string {
	lastDigitsArg := []byte("ss?c")

	if c.female {
		lastDigitsArg[2] = 'f'
	}

	if c.male {
		lastDigitsArg[2] = 'm'
	}

	return string(lastDigitsArg)
}

func main() {
	c := parseConfig()

	for range c.n {
		birthDate := getRandomBirthDate(getBirthRange(c))

		var s ssn.SSN
		s.SetDate(birthDate.In(time.Local))
		s.SetLastDigits(string(lastDigits(c)))

		fmt.Println(s.Format(true, true))
	}
}
