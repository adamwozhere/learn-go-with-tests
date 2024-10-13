package concurrency

import (
	"testing"
	"time"
)

// deliberately slow, waits 20 ms then returns true
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	// create a slice of 100 strings "a url"
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	// reset the timer before actually testing the implementation
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// check the websites url with the injected slowStubWebsiteChecker
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

// run `go test -bench=.`
// takes just over 2 seconds
