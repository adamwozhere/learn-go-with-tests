package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// it's important to be able to slice up requirements as small as possible
// to ensure we have working software - divide up our work and iterate on it:
// Print 3
// Print 3, 2, 1, Go!
// Wait a second between each line

// 1. first write a test to print "3"
// func TestCountdown(t *testing.T) {
// 	buffer := &bytes.Buffer{}

// 	Countdown(buffer)

// 	got := buffer.String()
// 	want := "3"
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// 2. Once that's working now test to print 3, 2, 1, Go!
// func TestCountdown(t *testing.T) {
// 	buffer := &bytes.Buffer{}

// 	Countdown(buffer)

// 	got := buffer.String()
// 	// using backticks will preserve space and lines,
// 	// e.g. this will output "3\n2\n1\nGo!"
// 	// (Althogh docs seem to show terminal printing on separate lines, but this just prints with `\n` characters, but the test works)
// 	want := `3
// 2
// 1
// Go!`

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// This all now functions correctly and tests pass, but the problem is that now the test takes 3 seconds to run!
// We have a dependency on `Sleep`, so if we can "Mock" it and use dependency injection then we can spy on the mock calls to test it

// func TestCountdown(t *testing.T) {
// 	buffer := &bytes.Buffer{}
// 	spySleeper := &SpySleeper{}

// 	// pass in the spy
// 	Countdown(buffer, spySleeper)

// 	got := buffer.String()
// 	want := `3
// 2
// 1
// Go!`

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}

// 	if spySleeper.Calls != 3 {
// 		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
// 	}
// }

// All now appears to be working with the sleeper, but we have the problem that the sleep calls could be out of order,
// e.g. a loop of 3 sleep calls then the loop of print statements, which is incorrect but the test will pass.

func TestCountdown(t *testing.T) {
	// now we have the `SpyCountdownOperactions` interface, we have to pass it in to Countdown,
	// but only need to check the buffer output to confirm the printed text is correct
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		// pass in the spy for both arguments
		Countdown(spySleepPrinter, spySleepPrinter)

		// now we can check on the order of operations (and amount)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

// Adding a test for the new configurable sleeper
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
