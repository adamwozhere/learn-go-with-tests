package racer

import (
	"fmt"
	"net/http"
	"time"
)

// // function with named return
// func Racer(a, b string) (winner string) {
// 	startA := time.Now()
// 	http.Get(a)
// 	aDuration := time.Since(startA)

// 	startB := time.Now()
// 	http.Get(b)
// 	bDuration := time.Since(startB)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// refactor with a helper function
// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// This is better but we should be using concurrency
// We don't care about the exact response times, we just want to know which one comes back first.
// To do this we can use `select` which helps sychronise processes easily.

// func Racer(a, b string) (winner string, error error) {
// 	// select:
// 	// recall from the previous chapter; you can wait for values to be sent to a channel with `myVar := <-ch`.
// 	// this is a *blocking* call as you're waiting for a value.
// 	// Select allows you to wait on *multiple* channels. The first one to send a value "wins", and the code
// 	// underneath the case is executed.

// 	// the ping function creates a channel for each url, whichever comes back first will be returned as the "winner"
// 	select {
// 	case <-ping(a):
// 		return a, nil
// 	case <-ping(b):
// 		return b, nil
// 		// `time.After` returns a `chan` and will send a signal down it after the amount of time you define.
// 		// perfect in this case, as it's possible that if you are listening to channels that never return a value, the code could block forever.
// 		// here, if neither return, then after 10 seconds the error will be returned.
// 	case <-time.After(10 * time.Second):
// 		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
// 	}
// }

// // creates a `chan struct{}` and returns it
// // we just want to signal we are done - closing the channel works perfectly
// // chan struct{} is the smallest data type available from a memory perspective,
// // we get no allocation versus a bool.
func ping(url string) chan struct{} {
	// we must always use the `make` function to create a channel;
	// rather than e.g. `var ch chan struct{}` - when you use the var keyword,
	// the variable will be initialised with the "zero" value of the type (string = "", int = 0 etc.)
	// for a channel this will be `nil` - which will block forever because you can't send (<-) `nil` to channels
	ch := make(chan struct{})
	// a goroutine that sends a signal into the channel once http.Get has completed
	go func() {
		// note that golangci-lint shows an error here in vscode:
		// Error return value of `http.Get` is not checked (errcheck)
		//nolint:errcheck // doesn't check for error, but we don't care about it in this case as we simply close the channel anyway.
		http.Get(url)
		close(ch)
	}()
	return ch
}

// The problem we have now is that the tests take 10 seconds to run. We should make the timeout configurable
// so that we can have a short timeout for tests, and a long one for the real world.

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout out waiting for %s and %s", a, b)
	}
}

// Summary
// Select helps you wait on multiple channels
// sometimes you'll want to include `time.After` in one of the cases to prevert the system from blocking forever
// `httptest` is a convenient way of creating test servers so you can have reliable, controllable tests -
// it uses the same interfaces as the "real" `net/http` servers which is consistent and less for us to learn.
