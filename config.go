package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/markysand/ssngenerator/v2/ssn"
	"github.com/rickb777/date"
)

type config struct {
	years, months, n                 int
	child, teen, adult, male, female bool
	format                           ssn.Format
	now                              date.Date
}

func parseConfig() *config {
	c := new(config)

	flag.IntVar(&c.years, "years", -1, "Set year")
	flag.IntVar(&c.months, "months", -1, "Set months")
	flag.IntVar(&c.n, "n", 1, "Number of ssn-s to generate")
	flag.BoolVar(&c.child, "child", false, "Child, 0-12 years")
	flag.BoolVar(&c.teen, "teen", false, "Teen, 13-17 years")
	flag.BoolVar(&c.adult, "adult", false, "Adult, 18-100 years")
	flag.BoolVar(&c.male, "male", false, "Generate male only")
	flag.BoolVar(&c.female, "female", false, "Generate female only")
	flag.Func("format", "Format (database, display, legacy)", func(s string) error {
		switch s {
		case "database", "":
			c.format = ssn.FormatDatabase
		case "display":
			c.format = ssn.FormatDisplay
		case "legacy":
			c.format = ssn.FormatLegacy
		default:
			return errors.New("invalid format")
		}
		return nil
	})

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

	c.now = date.Today()

	return c
}
