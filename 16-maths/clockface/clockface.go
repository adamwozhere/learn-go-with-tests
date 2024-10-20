package clockface

// notice this time this file is inside a `clockface` folder:
// modules and packages should reflect the folder structure -
// because the test is in a separate package (`clockface_test`),
// you have to import clockface which causes problems because the package is `clockface`,
// but it's inside the folder (and module) `16-maths`, so therefore you need to put the file
// inside of a `clockface` folder also in order that you can import `learn-go-with-tests/16-maths/clockface`.

import (
	"math"
	"time"
)

// we want to be able to create an analog clock SVG from a given time

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// ensure we don't use "magic numbers"
// const secondHandLength = 90
// const clockCentreX = 150
// const clockCentreY = 150

// // SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// // represented as a Point.
// func SecondHand(t time.Time) Point {
// 	// return Point{150, 60}
// 	// once the SecandHandPointTest is working we can work on this function
// 	// so that we can get the acceptance test passing

// 	p := secondHandPoint(t)
// 	// convert our unit vector to a point on the SVG:
// 	// scale it to the length of the hand,
// 	// flip it over the X axis to account for the SVG having an origin in the top left corner,
// 	// translate it to the right (so that it's coming from an orign of 150, 150)

// 	// not sure if this is redeclaring/making a new point each time
// 	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
// 	p = Point{p.X, -p.Y}                                      // flip
// 	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
// 	return p
// }

// after we have created the SVGWriter which is responsible for transforming the Point into the correct space and printing the xml
// we don't need the above `SecondHand` function anymore.

// add constants to get rid of the magic values
const (
	secondsInHalfClock = 30
	secondsInClick     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func secondsInRadians(t time.Time) float64 {
	// return math.Pi

	// return float64(t.Second()) * (math.Pi / 30)
	// floating point arithmetic is notoriously inaccurate, especially when we factor them up and down as here:
	// clockface_test.go:43: wanted 3.141592653589793 radians, but got 3.1415926535897936
	// by dividing `math.Pi` by 30 and then multiplying it by 30, we end up with a number thats no longer
	// the same as math.Pi.
	// we can refactor the equation by only dividing instead of dividing down them multiplying up.

	// return (math.Pi / (30 / (float64(t.Second()))))
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))

	// also note that computers don't like dividing by zero. In go it will give `+inf`
}

// func secondHandPoint(t time.Time) Point {
// 	angle := secondsInRadians(t)
// 	// use trigonometry to get the x, y values:
// 	// the hypotenuse is 1 (in our unit-circle, as radius is 1)
// 	x := math.Sin(angle)
// 	y := math.Cos(angle)

// 	return Point{x, y}
// }

func minutesInRadians(t time.Time) float64 {
	// we can just use the secondsInRadians function divided by 60:
	// fpr every second, the minute hand will move 1/60th of the angle the second hand moves.
	// Then we add on the movement for the minutes

	// return (secondsInRadians(t) / 60) +
	// 	(math.Pi / (30 / float64(t.Minute())))
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// func minuteHandPoint(t time.Time) Point {
// 	angle := minutesInRadians(t)
// 	x := math.Sin(angle)
// 	y := math.Cos(angle)

// 	return Point{x, y}
// }

// these Point functions are basically the same so lets refactor with a helper function
func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

// continue working to add the hour hand now
func hoursInRadians(t time.Time) float64 {
	// one full turn is an hour for the minute hand,
	// but for the hour hand its 12 hours, so divide the angle by 12,
	// then add the hour (note we need to module by 12, as the clock is not 24 hours!)

	// return (minutesInRadians(t) / 12) +
	// 	(math.Pi / (6 / float64(t.Hour()%12)))
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
