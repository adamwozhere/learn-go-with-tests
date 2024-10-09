package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Benchmarks can also be run in tests with a function beginning with `Benchmark`
// notice we use `*testing.B` and `b.N` in the loop, this means it will run `N` times.
// To run the benchmark use the CMD args: `go test -bench=.`
// it with print the results in the terminal:
// BenchmarkRepeat-8       12338925                95.62 ns/op
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("x", 3)
	fmt.Println(repeated)
	// Output: xxx
}
