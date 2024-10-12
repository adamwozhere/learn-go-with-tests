package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// the `Buffer` type from the `bytes` package implements the `Writer` interface
	// because it has the method `Write(p []byte) (n int, err error)`. So we can use this
	// to send in as our Writer for the for the greet function then test what was written to it.
	buffer := bytes.Buffer{}
	// here we are using the "address of" the buffer with the `&` symbol.
	// Not quite sure why we use address and not pointer.
	Greet(&buffer, "Adam")

	got := buffer.String()
	want := "Hello, Adam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
