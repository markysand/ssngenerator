package main

import (
	"io"
	"strings"
)

type Cutter struct {
	n         int
	basedepth int
	period    int
	output    io.Writer
}

func NewCutter(basedepth, period int, output io.Writer) *Cutter {
	return &Cutter{
		n:         0,
		basedepth: basedepth,
		period:    period,
		output:    output,
	}
}

func (cut *Cutter) Next() {
	cut.n++
	depth := getDepthOfPeriod(cut.n, cut.period) + cut.basedepth
	s := getSeparator(depth)
	cut.output.Write([]byte(s))
}

func (cut *Cutter) Reset() {
	cut.n = 0
}

func getDepthOfPeriod(index, period int) int {
	var n int
	for index != 0 {
		if index%period != 0 {
			break
		}
		index = index / period
		n++
	}
	return n
}

func getSeparator(depth int) string {
	var b strings.Builder
	switch {
	case depth == 1:
		b.WriteString("\t")
	case depth > 1:
		b.WriteString(strings.Repeat("\n", depth-1))
	}
	return b.String()
}
