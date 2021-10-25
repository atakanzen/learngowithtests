package store

import (
	"os"
)

// Tape is a struct that has a os.File field. This is for writing less data than current file data, or for such cases as editing or deleting data.
type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.File.Truncate(0)
	t.File.Seek(0, 0)
	return t.File.Write(p)
}
