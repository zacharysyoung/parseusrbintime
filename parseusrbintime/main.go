package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"

	record "github.com/zacharysyoung/parseusrbintime"
)

func main() {
	buf := &bytes.Buffer{}
	io.Copy(buf, os.Stdin)
	rec := record.New(buf.String())

	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"Real(s)", "User(s)", "Sys(s)", "Memory(B)"})
	writer.Write(rec.Strings())
	writer.Flush()
}
