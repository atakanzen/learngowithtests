package clock_test

import (
	"learngowithtests/maths/clock"
	"testing"
	"time"
)

/*
	every clock has a centre of (150, 150)
	the hour hand is 50 long
	the minute hand is 80 long
	the second hand is 90 long.
*/

func TestSecondHand(t *testing.T) {
	t.Run("at midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		want := clock.Point{X: 150, Y: 150 - 90}
		got := clock.SecondHand(tm)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("at 30 seconds", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

		want := clock.Point{X: 150, Y: 150 + 90}
		got := clock.SecondHand(tm)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
