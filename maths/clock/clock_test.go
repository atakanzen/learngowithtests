package clock_test

import (
	"learngowithtests/maths/clock"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{timeOnTheDate(0, 0, 30), math.Pi},
		{timeOnTheDate(0, 0, 0), 0},
		{timeOnTheDate(0, 0, 45), (math.Pi / 2) * 3},
		{timeOnTheDate(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.SecondsInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Errorf("got %v radians, want %v", got, c.angle)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clock.Point
	}{
		{timeOnTheDate(0, 0, 30), clock.Point{X: 0, Y: -1}},
		{timeOnTheDate(0, 0, 45), clock.Point{X: -1, Y: 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.SecondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("got %v vector point, want %v", got, c.point)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{timeOnTheDate(0, 30, 0), math.Pi},
		{timeOnTheDate(0, 0, 7), 7 * (math.Pi / (30 * 60))}, // find the angle of a second, then multiply it with 7
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.MinutesInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Errorf("got %v radians, want %v", got, c.angle)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clock.Point
	}{
		{timeOnTheDate(0, 30, 0), clock.Point{X: 0, Y: -1}},
		{timeOnTheDate(0, 45, 0), clock.Point{X: -1, Y: 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.MinuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("got %v vector point, want %v", got, c.point)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{timeOnTheDate(6, 0, 0), math.Pi},
		{timeOnTheDate(0, 0, 0), 0},
		{timeOnTheDate(21, 0, 0), math.Pi * 1.5},
		{timeOnTheDate(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.HoursInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Errorf("got %v radians, want %v", got, c.angle)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clock.Point
	}{
		{timeOnTheDate(6, 0, 0), clock.Point{X: 0, Y: -1}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.HourHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("got %v vector point, want %v", got, c.point)
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

func roughlyEqualPoint(a, b clock.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
