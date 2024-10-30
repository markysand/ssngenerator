package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/markysand/ssngenerator/v2/ssn"
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
	c := parseConfig()

	var gender ssn.Gender

	if c.male {
		gender = ssn.GenderMale
	} else if c.female {
		gender = ssn.GenderFemale
	}

	for range c.n {
		birthDate := getRandomBirthDate(getBirthRange(c))

		s := ssn.New(birthDate, gender)

		fmt.Println(s.Format(c.format))
	}
}
