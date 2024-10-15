package main1

import "sync"

// we want to make a counter that is safe to use concurrently

// type Counter struct {
// 	value int
// }

// func (c *Counter) Inc() {
// 	c.value++
// }

// func (c *Counter) Value() int {
// 	return c.value
// }

// As we know, it will fail if two goroutines try to write to the same bit of memory at the same time.
// A simple solution is to add a lock to the Counter - a `Mutex` (Mutual exclusion lock) provides this.
// The zero-value for a Mutex is an ulocked mutex.

type Counter struct {
	// Note that some examples will embed the mutex in the struct e.g.
	// sync.Mutex
	// then access it in the method e.g. c.Lock()
	// but this bad because methods of the embedded type become part of the public interface, which generally you don't want.
	// We don't want to expose `Lock` and `Unlock`
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// any goroutine calling Inc will acquire the lock on the Counter if they are first.
	// Any other goroutines will have to wait for it to be Unlocked before getting access.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

// As we must not copy the mutex (e.g. when sending into the assertCounter test function),
// we can provide a constuctor function which shows readers of the API that it's better not to initialise the type yourself.
func NewCounter() *Counter {
	return &Counter{}
}

// Summary:
// The `sync` package gives us:
// `Mutex` - allows us to add locks to our data
// `WaitGroup` - a means of waiting for goroutines to finish jobs
// It can be hard to know when to use channels or mutexes but in general:
// Use channels when passing ownership of data; use mutexes for managing state
// `go vet` can be used in build scripts to check for subtle bugs
// Don't use embedding because it's convenient - do you really want to expose the methods publicly? -
// It could be disastrous if some code was unlocking a mutex when it shouldn't be able to; it would be hard to track the bugs down.
