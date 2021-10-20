package main

import (
	"os"
)

// tape is a struct that has a os.File field. This is for writing less data than current file data, or for such cases as editing or deleting data.
type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
