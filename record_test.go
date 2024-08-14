package parseusrbintime

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	text := `
        1.01 real         0.99 user         0.02 sys
                   1  maximum resident set size
                   2  average shared memory size
                   3  average unshared data size
                   4  average unshared stack size
                   5  page reclaims
                   6  page faults
                   7  swaps
                   8  block input operations
                   9  block output operations
                  10  messages sent
                  11  messages received
                  12  signals received
                  13  voluntary context switches
                  14  involuntary context switches
                  15  instructions retired
                  16  cycles elapsed
                  17  peak memory footprint
`
	want := Record{
		Real: 1010 * time.Millisecond, User: 990 * time.Millisecond, Sys: 20 * time.Millisecond,

		Memory: 17,
	}

	if got := New(strings.TrimSpace(text)); !reflect.DeepEqual(got, want) {
		t.Errorf("\n  got %v\n want %v", got, want)
	}
}
