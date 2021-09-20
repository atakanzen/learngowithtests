package svg

import (
	"fmt"
	"io"
	"learngowithtests/maths/clock"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
	svgStart         = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

// secondHand writes the XML SVG line representation of the second hand of an anologue clock at the given time to the given writer
func secondHand(w io.Writer, t time.Time) {
	p := clock.SecondHandPoint(t)
	p = makeHand(p, secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

// minuteHand writes the XML SVG line representation of the minute hand of an analogue clock at the given time to the given writer
func minuteHand(w io.Writer, t time.Time) {
	p := clock.MinuteHandPoint(t)
	p = makeHand(p, minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

// hourHand writes the XML SVG line representation of the hour hand of an analogue clock at the given time to the given writer
func hourHand(w io.Writer, t time.Time) {
	p := clock.HourHandPoint(t)
	p = makeHand(p, hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p clock.Point, length float64) clock.Point {
	p = clock.Point{X: p.X * length, Y: p.Y * length}             // Scale to the length of the hand
	p = clock.Point{X: p.X, Y: -p.Y}                              // Flip over X axis to have an origin in the top left hand corner
	p = clock.Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY} // Translate to the right position (150,150)

	return p
}
