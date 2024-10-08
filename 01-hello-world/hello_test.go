package main

import "testing"

// Testing in GO
// cmd `go test` to run the tests, ensure a module (go.mod) is present: cmd `go mod init <module_name>`
// the file must be named like `xxx_test.go`
// the test function must start with the word `Test`
// the test function takes one argument only e.g. (t *testing.T) imported from "testing"
// for now, just know that `t` of type `*testing.T` is a "hook" infor the testing framework
// running `go test` will produce e.g.
// PASS
// ok      hello   0.121s

// create a test case first before the actal function (Test Driven Development)
// Writing a test first and seeing it fail, lets us know that we have written a relevant test
// for our requirements, and seen that it produces an easy to understand description of the failure.
// we then write the smallest amount of code to make it pass so we know we have working software.
// then we can refactor, knowing that everything is tested and that we have good code that is easy to work with
func TestHello(t *testing.T) {
	// create subtests with `t.Run` similar to jest etc.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Adam", "")
		want := "Hello, Adam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jacques", "French")
		want := "Bonjour, Jacques"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Hans", "German")
		want := "Hallo, Hans"
		assertCorrectMessage(t, got, want)
	})
}

// use interface testing.TB which satisfies T and B so you can use helper functions from Testing and Benchmarking
func assertCorrectMessage(t testing.TB, got, want string) {
	// tells the suite that this is a helper function,
	// if tests fail then the line number of the `test` function will be reported instead of this helper function
	t.Helper()
	if got != want {
		// `Errorf` prints a message and fails the test
		// the `f` stands for formatting and allows building a string with placeholder values `%q` which get wrapped in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
