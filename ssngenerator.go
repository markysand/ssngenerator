package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/markysand/ssn"
)

const hoursPerYear = 24 * 365.25

func main() {
	flagSafe := flag.Bool("safe", false, "Safe flag will restrict to safe 980C-999C three last digits SSN")
	flagMin := flag.Float64("min", 0, "Sets min age boundary")
	flagMax := flag.Float64("max", 100, "Sets max age boundary")
	flagCheck := flag.Bool("check", false, "Will check an SSN, and return correct checksum if needed")
	flagMale := flag.Bool("male", false, "Generate male only")
	flagFemale := flag.Bool("female", false, "Generate female only")
	flagNo := flag.Int("n", 1, "Select number of SSNs to generate")
	flag.Usage = func() {
		fmt.Println("SSN-Generator is a tool for swedish social security numbers - SSN-s. SSN-s that end with 98xx or 99xx are not used by real living persons and are considered safe for testing. -check and -safe can be combined - will also return the SSN provided as a safe variant")
		flag.CommandLine.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	l := len(args)
	switch {
	case *flagCheck && l >= 1:
		input := args[0]
		n, err := ssn.NewSSNFromString(input)
		switch err {
		case nil:
			fmt.Printf("%v is a valid SSN\n", n)
			fmt.Printf("Age: %v, Gender: %v\n",
				math.Round(n.Age(time.Now()).Hours()/hoursPerYear*100)/100,
				func() string {
					if n.Female() {
						return "FEMALE"
					}
					return "MALE"
				}(),
			)
			if *flagSafe {
				n.SetLastDigits("ss*c")
				fmt.Printf("As a safe ssn, it should be %v\n", n)
			}
		case ssn.ErrChecksum:
			n.SetLastDigits("***c")
			fmt.Printf("%v has incorrect checksum - it should be: %v\n", input, n)
		default:
			fmt.Printf("%v has invalid format - %v\n", input, err.Error())
		}
	case l >= 1:
		fmt.Println("Not sure what you want me to do... maybe you want me to check an ssn - in that case use the 'check' flag")
		flag.PrintDefaults()
	default:
		last := []byte("???c")
		if *flagSafe {
			last[0], last[1] = byte('s'), byte('s')
		}
		if *flagMale {
			last[2] = byte('m')
		} else if *flagFemale {
			last[2] = byte('f')
		}
		max := time.Duration(*flagMax*hoursPerYear) * time.Hour
		min := time.Duration(*flagMin*hoursPerYear) * time.Hour
		separator := NewCutter(1, 5, os.Stdout)
		for i := 0; i < *flagNo; i++ {
			var n ssn.SSN
			n.SetDate(ssn.GetRandomTime(max, min))
			n.SetLastDigits(string(last))
			fmt.Print(n)
			separator.Next()
		}
	}
}
