package ssn

import (
	"math/rand"
	"strings"

	"github.com/rickb777/date"
)

const (
	indexYear1 int = iota
	indexYear2
	indexYear3
	indexYear4
	indexMonth1
	indexMonth2
	indexDay1
	indexDay2
	indexRegion1
	indexRegion2
	indexGender
	indexCheckSum
)

type SSN [12]int

const zeroChar = '0'

func doubled(n int) int {
	value := n * 2

	if value > 9 {
		return value - 9
	}

	return value
}

func (n SSN) GetCheckSum() int {
	var sum int
	for i := 2; i < 11; i++ {
		if i%2 == 0 {
			sum += doubled(n[i])
		} else {
			sum += n[i]
		}
	}

	result := (10 - sum%10) % 10

	return result
}

func (n *SSN) SetChecksum() {
	(*n)[indexCheckSum] = n.GetCheckSum()
}

type Format int

const (
	// Century digits and no separator
	FormatDatabase = iota
	// Century digits and separator, but no + since it is redundant
	FormatDisplay
	// No century digits, original separator that turns to + the year the person will be 100 years old
	FormatLegacy
)

func (f Format) ShowCentury() bool {
	switch f {
	case FormatDatabase, FormatDisplay:
		return true
	default:
		return false
	}
}

func (f Format) ShowDash() bool {
	switch f {
	case FormatDisplay, FormatLegacy:
		return true
	default:
		return false
	}
}

func (n *SSN) Year() int {
	return n[indexYear1]*1000 + n[indexYear2]*100 + n[indexYear3]*10 + n[indexYear4]
}

func (n *SSN) Format(f Format) string {
	var b strings.Builder

	for index, v := range n {
		if !f.ShowCentury() && (index == indexYear1 || index == indexYear2) {
			continue
		}

		b.WriteByte(byte(v + zeroChar))

		if index == 7 && f.ShowDash() {
			// Use `+` for 100 year olds when century digits are missing
			if f == FormatLegacy && date.Today().Year()-n.Year() >= 100 {
				b.WriteByte('+')
			} else {
				b.WriteByte('-')
			}

		}
	}

	return b.String()
}

func (n *SSN) SetBirthDate(year int, month int, day int) {
	(*n)[indexYear1] = year / 1000
	(*n)[indexYear2] = year / 100 % 10
	(*n)[indexYear3] = year / 10 % 10
	(*n)[indexYear4] = year % 10

	(*n)[indexMonth1] = month / 10
	(*n)[indexMonth2] = month % 10

	(*n)[indexDay1] = day / 10
	(*n)[indexDay2] = day % 10
}

type Gender int

const (
	GenderUnknown = iota
	GenderMale
	GenderFemale
)

func (n *SSN) SetGender(g Gender) {
	switch g {
	case GenderMale:
		(*n)[indexGender] = 1 + rand.Intn(5)*2
	case GenderFemale:
		(*n)[indexGender] = 0 + rand.Intn(5)*2
	default:
		(*n)[indexGender] = rand.Intn(10)
	}
}

func (n *SSN) SetRegionDigits() {
	(*n)[indexRegion1] = 9
	(*n)[indexRegion2] = 8 + rand.Intn(2)

}

func New(d date.Date, g Gender) *SSN {
	var ssn SSN

	ssn.SetBirthDate(d.Year(), int(d.Month()), d.Day())
	ssn.SetGender(g)
	ssn.SetRegionDigits()
	ssn.SetChecksum()

	return &ssn
}
