package contexts

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// type SpyStore struct {
// 	response  string
// 	cancelled bool
// 	t         *testing.T
// }

// func (s *SpyStore) Fetch() string {
// 	time.Sleep(100 * time.Millisecond)
// 	return s.response
// }

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// refactor by adding assertion methods on the spy
// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }

// func TestServer(t *testing.T) {
// 	data := "hello, world"

// 	t.Run("returns data from the store", func(t *testing.T) {
// 		// remember to add `t` when creating the spy.
// 		store := &SpyStore{response: data, t: t}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		if response.Body.String() != data {
// 			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
// 		}

// 		// if store.cancelled {
// 		// 	t.Error("it should not have cancelled the store")
// 		// }

// 		// use assertion method instead:
// 		store.assertWasNotCancelled()
// 	})

// 	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
// 		store := &SpyStore{response: data, t: t}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)

// 		// the context package provides functions to derive new Context values from existing ones.
// 		// These values form a tree: when a Context is cancelled, all Contexts derived from it are also cancelled.
// 		cancellingCtx, cancel := context.WithCancel(request.Context())
// 		time.AfterFunc(5*time.Millisecond, cancel)
// 		request = request.WithContext(cancellingCtx)

// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		// if !store.cancelled {
// 		// 	t.Error("store was not told to cancel")
// 		// }

// 		// use assertion method instead
// store.assertWasCancelled()
// 	})
// }

// refactoring tests so that the context can be passed to the Store
// change the SpyStore
type SpyStore struct {
	response string
	t        *testing.T
}

// we have to make the spy act like a real method that works with context.
// we are simulating a slow process where we build the result slowly by appending the string,
// character by character in a gorouting. When the goroutine finishes, it write the string to the data channel.
// The goroutine listens for the `ctx.Done` and will stop the work if a signal is sent in that channel.
// Finally the code uses another `select` to wait for that goroutine to finish its work or for the cancellation to occur.

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// we need to test that we do not write any kind of response on the error case.
// `httptest.ResponseRecorder` doesn't have a way of figuring this out wo we'll have to write our own spy.
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

// update the tests
func TestContext(t *testing.T) {
	data := "hello world"

	t.Run("returns data fromt he store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells the store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
