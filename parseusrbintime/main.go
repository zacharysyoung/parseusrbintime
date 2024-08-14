package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	record "github.com/zacharysyoung/parseusrbintime"
)

func main() {
	buf := &bytes.Buffer{}
	io.Copy(buf, os.Stdin)
	rec := record.New(buf.String())

	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"Real(ms)", "User(ms)", "Sys(ms)", "Memory(B)"})
	writer.Write(strings(rec))
	writer.Flush()
}

// dtoa returns a normalized string of d in milliseconds.
func dtoa(d time.Duration) string { return strconv.Itoa(int(d.Milliseconds())) }
func itoa(i int) string           { return strconv.Itoa(i) }

// Strings returns a slice of strings for all values of the record.
func strings(r record.Record) []string {
	return []string{
		dtoa(r.Real), dtoa(r.User), dtoa(r.Sys),

		itoa(r.Memory),
	}
}
