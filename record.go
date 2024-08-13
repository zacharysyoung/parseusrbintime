package parseusrbintime

import (
	"strconv"
	"strings"
)

// Record represents the real, user, sys timing, and the peak
// memory footprint values from the output of `/usr/bin/time -l ...`.
type Record struct {
	Real, User, Sys float64

	Memory int // peak memory footprint
}

// New parses text that looks like the example above.
//
// The text arg must look like:
//
//	1.23 real         4.56 user         7.89 sys
//	          ...  ...
//	         9876  peak memory footprint
//
// with no empty lines before the timings line.  The memory
// footprint line must be the 18th line in text.
func New(text string) Record {
	lines := strings.Split(text, "\n")

	fields := strings.Fields(lines[0])

	return Record{
		Real: atof(fields[0]),
		User: atof(fields[2]),
		Sys:  atof(fields[4]),

		Memory: linetoi(lines[17]),
	}
}

func ftoa(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }
func itoa(i int) string     { return strconv.Itoa(i) }

// Strings returns a slice of strings for all values of the record.
func (r Record) Strings() []string {
	return []string{
		ftoa(r.Real), ftoa(r.User), ftoa(r.Sys),

		itoa(r.Memory),
	}
}

func atof(s string) (f float64) { f, _ = strconv.ParseFloat(s, 64); return }
func atoi(s string) (i int)     { i, _ = strconv.Atoi(s); return }

// linetoi returns the int component of a line like,
// "   123  peak memory footprint".
func linetoi(line string) int { return atoi(strings.Fields(line)[0]) }
