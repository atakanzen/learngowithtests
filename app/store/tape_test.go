package store_test

import (
	"io/ioutil"
	"learngowithtests/app/helper"
	"learngowithtests/app/store"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := helper.CreateTmpFile(t, "12345")
	defer clean()

	tape := &store.Tape{file}
	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
