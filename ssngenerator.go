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

// norwegian ssn checksum
func checksum(ssn string) int {
	weights := []int{3, 7, 6, 1, 8, 9, 4, 5, 2}
	sum := 0

	for i, r := range ssn[:len(ssn)-1] {
		sum += int(r-'0') * weights[i]
	}

	return (11 - sum%11) % 11
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
