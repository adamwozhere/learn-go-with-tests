package clockface

import (
	"math"
	"testing"
	"time"
)

// we can use trigonometry to work out the angles and position of the hands
// - using the math library we will need to use radians instead of degrees.
// lets create separate tests to test this, outside of the acceptance test.

// these tests could be moved out or even deleted later - we're using them to test
// that we're getting the maths correct for working out the angles etc.

// func TestSecondsInRadians(t *testing.T) {
// 	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
// 	// a full turn of a circle is 2*Pi radians, so we know that 30 seconds should just be Pi.
// 	want := math.Pi
// 	got := secondsInRadians(thirtySeconds)

// 	if want != got {
// 		t.Fatalf("wanted %v radians, but got %v", want, got)
// 	}
// }

// extend the test as a table to test different seconds angles
func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

// tests for returning the Point of the second hand based on a unit circle
// (where the radius of the circle is 1):
// zero would be the centre of the circle with axis X going from -1 on the left to 1 on the right
// and Y axis going from -1 on the bottom to 1 on the top.
func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			// this still isn't completely accurate,
			// we could increase accuraticy with rational type `Rat` from `math/big` package
			// but as these numbers are infinitesimal way down to 16th decimal place, we can say they're roughly equal
			// and as we are just creating an SVG from this it won't matter.
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("wanted %v point, but got %v", c.point, got)
			}
		})
	}
}

// once we have the acceptance test working for the SVGWriter for the second hand, we can work on the minute hand
func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		// 7 seconds past the hour so the the minute hand should be over 12 and a little bit
		// (60 seconds in a minute, 30 mins in half a turn of the circle: math.Pi radians)
		// so 30 * 60 seconds in a half turn - so it should be at 7 * (math.Pir / ( 30 * 60)) radians
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		// again the hour hand should turn a bit more based on seconds and minutes
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

// hepler functions

// returns a date from hours, minutes, seconds
func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	// returns a string format according to the layout - makes it easier to read in the tests
	return t.Format("15:04:05")
}

// equality helpers
func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
