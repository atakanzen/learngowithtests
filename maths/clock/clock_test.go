package clock

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{time: timeOnTheDate(0, 0, 30), angle: math.Pi},
		{time: timeOnTheDate(0, 0, 0), angle: 0},
		{time: timeOnTheDate(0, 0, 45), angle: (math.Pi / 2) * 3},
		{time: timeOnTheDate(0, 0, 7), angle: (math.Pi / 30) * 7},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondsInRadians(test.time)
			if got != test.angle {
				t.Errorf("got %v radians, want %v", got, test.angle)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{time: timeOnTheDate(0, 0, 30), point: Point{X: 0, Y: -1}},
		{time: timeOnTheDate(0, 0, 45), point: Point{X: -1, Y: 0}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v vector, want %v", got, test.point)
			}
		})
	}
}

func timeOnTheDate(hours, minutes, seconds int) time.Time {
	return time.Date(2017, time.October, 22, hours, minutes, seconds, 0, time.UTC)
}

func testName(time time.Time) string {
	return time.Format("22:10:17")
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
