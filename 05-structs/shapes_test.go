package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// `%.2f` means print to 2 decimal places (f means float)
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	t.Run("rectangles", func(t *testing.T) {

// 		rectangle := Rectangle{12.0, 6.0}
// 		got := rectangle.Area()
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			// using `%g` will print a more precise decimal number than `%f`,
// 			// a radius of 1.5 would print 7.068583 with `f`
// 			// whereas `g` would print 7.0685834705770345
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }

// refactoring the Area tests
// All we want to do is to take a collection of "shapes", call the Area() method
// on them and check the result. We want a checkArea function that can be passed
// both Rectangles and Circles - we can do this with `interfaces`.

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// further refactoring with "table driven tests"

// func TestArea(t *testing.T) {
// 	// here we create an "anonymous struct", areaTests.
// 	// We declare a slice of structs using `[]struct` with two fields,
// 	// then fill the slice with cases - allowing us to loop over them.
// 	// Be sure that table driven tests are needed before using them,
// 	// but they can be great when testing various implementations of an interface,
// 	// or if the data being passed has lots of different requirements that need testing.
// 	// It's now very easy to add a new test case for a Shape e.g. Triangle
// 	areaTests := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{Rectangle{12, 6}, 72.0},
// 		{Circle{10}, 314.1592653589793},
// 		{Triangle{12, 6}, 36.0},
// 	}

// 	for _, tt := range areaTests {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %g want %g", got, tt.want)
// 		}
// 	}
// }

// Making the table driven test easier to read
// using named fields in the structs, and "hasArea" instead of "want",
// using a name in the table, with each test in `t.Run` to use that name,
// and printing the struct values with `%#v`.
// When erroring, these changes will make it more descriptive.
// e.g. `shapes_test.go:119: main.Triangle{Base:12, Height:6} got 36 want 37`.
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// `%#v` will print the struc values, so we can see what exactly is being input
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
