package contexts

import (
	"context"
	"fmt"
	"net/http"
)

// software often kicks off long-running, resource intensive processes (often goroutines).
// If the action that caused this gets cancelled or fails for some reason, you need to stop
// these proceses n a consistent way through your application.
// the `context` package helps us to manage such long-running processes.
// In this example we will setup a server with a scenario where a user cancels the request
// before the data can be retrieved and make sure the process is told to stop.

// type Store interface {
// 	Fetch() string
// 	Cancel()
// }

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		store.Cancel()
// 		fmt.Fprint(w, store.Fetch())
// 	}

// 	// added `store.Cancel()` which makes the cancel test work,
// 	// but we shouldn't be cancelling Store before we fetch on *every request* !
// }

// A more sensible implementation
// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()

// 		data := make(chan string, 1)

// 		// context's method `Done` returns a channel which is sent a signal when
// 		// the context is "done" or "cancelled". We listen to that signal and call `store.Cancel`
// 		// if get it but we want to ignore it if our Store manages to Fetch before it.

// 		// To manage this we run Fetch in a goroutine and it will write the result in a a new channel `data`.
// 		// we then use `select` to effectively race the two asynchronous processes and then we either write a response
// 		// or `Cancel`.

// 		go func() {
// 			data <- store.Fetch()
// 		}()

// 		select {
// 		case d := <-data:
// 			fmt.Fprint(w, d)
// 		case <-ctx.Done():
// 			store.Cancel()
// 		}
// 	}
// }

// this doesn't feel good that the severs has to be concerned with manually cancelling `Store`.
// What if Store also happens to depend on other slow-running processes? We'll have to make sure that Store.Cancel
// correctly propagates the cancellation to all it's dependants.
// One of the main points of Context isthat it's a consistent way of offering cancellation.
// From the go doc: Incoming server requests should create a Context, and outgoing calls to servers should accept a Context.
// The chain of function calls between them must propagate the Context, optionally replacing it with a derived context, e.g. created using WithCancel etc.

// So lets try and pass through the context to the Store and let it be responsible, that way it can also pass the context through to it's dependants,
// and they too can be responsible for stopping themselves.

// update the Store interface
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}

// Summary:
// use Context to manage cancellation
// write functions taht accept `context` and use it to cancel itself by using goroutines, `select` and channels.
// It is advised not to use `context.Value` even though it feels convenient, as it's just an untyped map so there is no type safety.
// -- in short, if a function needs some values, pass them as typed parameters rather than try to fetch them from `context.Value` if possible.
// This ensures they are statically checked and documented for people to see. However, be aware there may be some situations where it is useful,
// e.g. including a trace id? - not every function would need it, and so otherwise it would make the function signatures messy.
// (Refer to the Go blog documention which has advice on this https://go.dev/blog/context).
