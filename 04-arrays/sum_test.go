package main

// notice we are using package main this time,
// tests will not work because you cannot `go mod init main`.
// According to common practice, package main will only contain integration
// or other packages and not unit-testable code
// and hence Go will not allow you to import a package with name `main` -
// therefore rename the main module to something else e.g. `go mod init arrays`.

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// We had two tests, for 5 numbers and 3 numbers,
	// but we don't need both as if it works on a slice of 5 numbers it should work on a slice of N numbers.
	// Removing one of the tests we can check coverage by using command `go test -cover`
	// which will show: `coverage: 100.0% of statements`

	t.Run("collection of 5 numbers", func(t *testing.T) {
		// using a Slice
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			// `%d` prints the number in base 10
			// `%v` prints the value in a default number (good here as we can print the input array)
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

// func TestSumAll(t *testing.T) {
// 	got := SumAll([]int{1, 2}, []int{0, 9})
// 	want := []int{3, 9}

// 	// when the initial Step 1. function returns nil, the compiler errors:
// 	// invalid operation: got != want (slice can only be compared to nil).
// 	// Go does not let you use equality operators with slices.
// 	// You could write a function to iterate each slice, but we can use `reflect.DeepEqual` instead.

// 	// if got != want {
// 	// 	t.Errorf("got %v want %v", got, want)
// 	// }

// 	// now the test will compile and run because we are using `DeepEqual`,
// 	// which is useful for seeing if ANY two variables are the same.
// 	// Note that `DeepEqual` is NOT type safe. The code will compile if we changed `got` to a string etc.

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}

// 	// Grom Go 1.21 slices standard packages has `slices.Equal` to do a simple shallow compare on slices,
// 	// where you don't need to worry about the types. Note that this function expects elements to be comparable,
// 	// so it can't be applied to slices with non-comparable elements like 2D slices.
// }

func TestSumAllTails(t *testing.T) {
	// reducing the repeated code, this time by assigning a function to a variable,
	// this technique can be useful when you want to bind a function to other local variables in scope.
	// it also allows you to reduce the surface area of your API, as it cannot be used outside of this test.
	// It gives type safety too as you cannot create a new test with checkSums
	// Hiding variables and functions that don't need to be exported is an important design consideration.
	checkSums := func(t testing.TB, got, want []int) {
		// here I am not sure why `t testing.TB` does not include the asterisk,
		// I think because `t` is being passed in which is already in scope as a pointer
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	// initially this test will compile but have a runtime error, as index is out of bounds,
	// because we are trying to access index 1 on and empty slice (with [1:])
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
