package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Testable Examples can be provided which will be present in godocs,
// They live inside the xxx_test.go files and begin with Example in the function name
// these are good as often a readme file etc. may become out of date, unchecked and incorrect.
// Testable Examples are compiled whenever tests are excecuted,
// this means documentation will always reflect the current code behavior
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// notice the comment inside the ExampleAdd function `Output: 6`,
// if this is removed, the test will not be executed

// running `godoc -http=localhost:6060` will launch the docs,
// visit localhost:6060/pkg/integers and you can see the Add function,
// it will show an Example with the code and `Output: 6` under it!

// TLDR: adding a comment above a function will describe it in the docs.
// An Example can be added to the docs by including an ExampleFunction in the test file
// with an `Output: value` comment inside it.
