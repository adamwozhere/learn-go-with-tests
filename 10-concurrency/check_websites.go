package concurrency

// dependency that we inject so that we can test and run it in production.
// Still a bit unsure why it isn't an interface?
type WebsiteChecker func(string) bool

// Note: actual function to check real website is not implemented in this exercise, only the testing function.

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		results[url] = wc(url)
// 	}

// 	return results
// }

// Concurrency--is having more than one thing in progress.
// instead of checking each url one after another, we can use `goroutines` to create another process.
// once the loop gets to a url to check, it won't block, it will simply start a process -
// but not wait for it to finish - and continue looping creating processes for each item in the loop.
// To start a new goroutine we put `go` in front of a function.
// For this case we need to use an anonymous function which we can immediately call by putting braces after it.

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		// wrapped in an anoynmous function that we can use as a goroutine
// 		// anonymous functions keep lexical sccope so we can add to the `results` map

// 		// this will only put one url into the map as each goroutine iteration will have
// 		// a reference to the `url` variable, they don't have their own independent copy,
// 		// so they are ALL writing the value that `url` has at the end of the iteration - the last url.
// 		// go func() {
// 		// 	results[url] = wc(url)
// 		// }()

// 		// we can pass in the current url value to the goroutines so they have their own independent copy
// 		go func(url string) {
// 			results[url] = wc(url)
// 		}(url)
// 	}
// 	// the function could keep going and return an empty map as the goroutines may not have finished yet
// 	// Could try waiting:
// 	// time.Sleep(2 * time.Second)
// 	// but this causes `fatal error: concurrent map writes` - sometimes two goroutines can write to the map
// 	// at exactly the same time. Maps in Go don't like this so will cause a fatal error as it's a "race condition".
// 	// You can run `go test -race` to see the results. (Note: cannot get it to work on my machine as needs cgo setting which doesn't work)
// 	// - it will give a "warning: data race" and will probably indicate that two goroutines are trying to write to the same block of
// 	// memory at the same time. To solve this we need to use `Channels`.

// 	// Channels are a Go data structure that can both receive and send values - they allow communication between different processes.

// 	return results
// }

// Rewriting the function using channels

// a struct for result where the values are anonymous
// a bit like a tuple in JS ?
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// use make to create a channel of type result
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// use the "send statement" (`<-`) instead of writing to the map directly
			// A send statement uses the `<-` operator with the channel on the left and a value on the right.
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// this next part is a loop of "receive expressions" which assigns a value recieved from a channel to a variable.
	for i := 0; i < len(urls); i++ {
		// receive expression - receives the channel (on the right of the `<-`)
		r := <-resultChannel
		// we then use the result received to update the map.
		results[r.string] = r.bool
	}

	return results
}

// when running bench again with channels and goroutines, it will take nanoseconds instead.

// SUMMARY
// We use goroutines each with their own reference to a url, to start a process to check the url concurrently.
// To avoid race conditions we send the result of checking a url to a channel.
// The channels are then used in a second loop with the recieve expression to linearly write the results to the map.
// This means that the part we wanted to speed up haves concurrently, but the part that cannot happen simultaneously (writing to the map),
// still happens linearly. (NOTE: I'm unsure how the second loop knows that the channel has finished getting all the data)
