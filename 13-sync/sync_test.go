package main1

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		// counter := Counter{}
		// use the NewCounter constructor function instead - gives us a pointer to use in assertCounter
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// `syc.WaitGroup` is a way of synchronising concurrent processes.
		// It waits for a collection of goroutines to finish.
		var wg sync.WaitGroup
		// The main goroutine calls `Add` to set the number of goroutines to wait for
		wg.Add(wantedCount)

		// then each goroutine runs and calls `Done` when finished.
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		// at the same time Wait can be used to block until all goroutines have finished
		wg.Wait()

		// once `wg.Wait()` has finished we can assert the values
		assertCounter(t, counter, wantedCount)
	})
}

// make a helper as we'll be testing more around the Counter
// we get a "copylocks" warning - more info can be seen in command `go vet` -
// the mutex documentation states that a mutex must not be copied after first use.
// - we therefore need to pass a pointer to our Counter instead
func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
