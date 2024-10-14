package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func TestRacer(t *testing.T) {

// 	// instead of testing against real urls, use the `httptest` package.
// 	// `httptest.NewServer` takes an `http.HandlerFunc` which we are sending in via an anonymous function.
// 	// http.HandlerFunc is a type that looks like this: `type HandlerFunc func(ResponseWriter, *Request)`.
// 	// This essentially how you would write a server in Go, but we wrap it in `httptest.NewServer` to make it easier to use in testing.
// 	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(20 * time.Millisecond)
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	slowURL := slowServer.URL
// 	fastURL := fastServer.URL

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}

// 	slowServer.Close()
// 	fastServer.Close()
// }

// the test has a lot of repeated code, so refactor with a helper function

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		// `defer` keyword used with a function -  will call that function at the end of the containing function.
		// this allows us to keep the server code together so it's easier to read.
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		// check for error first, as if there is an error the test should not continue,
		// (stop immediately with `t.Fatalf`)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		// just use the same server for this test
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		// use ConfigurableRacer with a very short timout for this test
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't ge one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
