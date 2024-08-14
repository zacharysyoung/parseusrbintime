package parseusrbintime

import (
	"strconv"
	"strings"
	"time"
)

// Record represents the real, user, sys timing, and the peak
// memory footprint values from the output of `/usr/bin/time -l ...`.
type Record struct {
	Real, User, Sys time.Duration

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
		Real: atod(fields[0]),
		User: atod(fields[2]),
		Sys:  atod(fields[4]),

		Memory: linetoi(lines[17]),
	}
}

// dtoa returns a normalized string of d in milliseconds.
func dtoa(d time.Duration) string { return strconv.Itoa(int(d.Milliseconds())) }
func itoa(i int) string           { return strconv.Itoa(i) }

// Strings returns a slice of strings for all values of the record.
func (r Record) Strings() []string {
	return []string{
		dtoa(r.Real), dtoa(r.User), dtoa(r.Sys),

		itoa(r.Memory),
	}
}

// atod converts a seconds value with a decimal to a time.Duration.
func atod(s string) time.Duration { d, _ := time.ParseDuration(s + "s"); return d }
func atoi(s string) int           { i, _ := strconv.Atoi(s); return i }

// linetoi returns the int component of a line like,
// "   123  peak memory footprint".
func linetoi(line string) int { return atoi(strings.Fields(line)[0]) }
