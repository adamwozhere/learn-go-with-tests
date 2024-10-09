package main

// notice we are using package main this time,
// tests will not work because you cannot `go mod init main`.
// According to common practice, package main will only contain integration
// or other packages and not unit-testable code
// and hence Go will not allow you to import a package with name `main` -
// therefore rename the main module to something else e.g. `go mod init arrays`.

import "testing"

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
