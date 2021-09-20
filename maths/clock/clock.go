// Package clock provides functions that calculate the positions of the hands.
// of an analogue clock.
package clock

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point is a two dimensional Cartesian coordinate, it represents the unit vector from the origin of a clock hand
type Point struct {
	X float64
	Y float64
}

// SecondsInRadians returns the angle of second hand from 12 o'clock in radians
func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

// SecondHandPoint returns the unit vector of the second hand at time `t` as a Point
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

// MinutesInRadians returns the angle of minute hand from 12 o'clock in radians
func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// MinuteHandPoint returns the unit vector of the minute hand at time `t` as a Point
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

// HoursInRadians returns the angle of hour hand from 12 o'clock in radians
func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

// HourHandPoint returns the unit vector of the hour hand at time `t` as a Point
func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}
