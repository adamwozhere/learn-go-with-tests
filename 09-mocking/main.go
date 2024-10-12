package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// write a function that will count down a 1 second pause in between:
// 3
// 2
// 1
// Go!

// When writing the minimal function to make the test work:
// func Countdown() {}
// the test will show us that we need `*bytes.Buffer`:
// too many arguments in call to Countdown
//         have (*bytes.Buffer)
//         want ()

// func Countdown(out *bytes.Buffer) {}
// Now the test runs:
// countdown_test.go:22: got "" want "3"

// 2. Write enough code to make it pass
// func Countdown(out *bytes.Buffer) {
// 	fmt.Fprint(out, "3")
// }
// this passes the first stage of the test, to just print "3"

// 3. Refactor - `*bytes.Buffer` works but we should be using a general purpose interface instead:
// func Countdown(out io.Writer) {
// 	fmt.Fprint(out, "3")
// }

// create a main function to run the countdown so we can see that we have some working software
// `go run main.go`
// func main() {
// 	Countdown(os.Stdout)
// }

// 4. Write enough code to make the 3, 2, 1, Go! test pass
// func Countdown(out io.Writer) {
// 	for i := 3; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}
// 	fmt.Fprint(out, "Go!")
// }

// 5. Refactor it with constants, and with the 1 second delay
const finalWord = "Go!"
const countdownStart = 3

// func Countdown(out io.Writer) {
// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 		time.Sleep(1 * time.Second)
// 	}
// 	fmt.Fprint(out, finalWord)
// }

// This all now functions correctly and tests pass, but the problem is that now the test takes 3 seconds to run!
// We have a dependency on `Sleep`, so if we can "Mock" it and use dependency injection then we can spy on the mock calls to test it

// define the dependency as an interface - this lets us use a *real* Sleeper in `main` and a *spy sleeper* in tests
type Sleeper interface {
	Sleep()
}

// Now we need to mack a *mock* of it for the tests to use
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// "spies" are a kind mock that can record how a dependencies is used. They can record the arguments sent in,
// how many times it's been called etc. In this case we're keeping track of many times Sleep() is called

// Now write the rewrite the function to accept Sleeper
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		// we call sleeper instead of directly `time.Sleep(1 * time.Second)`
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

// create a "real" sleeper which implements the interface we need
type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// now we can call it properly with out main() function
// func main() {
// 	// call countdown with the default sleeper
// 	sleeper := &DefaultSleeper{}
// 	Countdown(os.Stdout, sleeper)
// }

// We need to fix the problem for testing, that the order of the sleep and the prints are correct,
// so as we have two dependencies we need to create a spy that spies on both
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Now we have a spy that implements both `io.Writer` and `Sleeper`, recording every call into one slice.

// Extending the Sleeper to be configurable

// create a new type that accepts what we need for configuration and testing
type ConfigurableSleeper struct {
	duration time.Duration
	// sleep signature is the same as the interface, so we can also spy on it in tests
	// not sure why this is defined *inside* the struct this time - seems a bit like an interface?
	sleep func(time.Duration)
}

// create the spy for tests
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// add the sleep function for configurable sleeper,
// which takes the duration value from the struct and calls sleep which is used for the tests
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Now we can use the configurable sleeping the main() function
// as it's now using configurable sleeper, the defaultSleeper is no longer needed and can be deleted.
// this now means that we have a more generic sleeper with arbitrarily long countdowns
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

// Use spies with caution as it lets you see inside the algorythm you are writing, which can be useful
// but that means tighter coupling between you test code and implementaition.
// Be sure you actually care about these details if you're going to spy on them.
