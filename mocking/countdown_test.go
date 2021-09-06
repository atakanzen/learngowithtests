package mocking_test

import (
	"bytes"
	"learngowithtests/mocking"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &mocking.SpySleeper{}

	mocking.Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
