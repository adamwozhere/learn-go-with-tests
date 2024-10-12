package main

import (
	"fmt"
	"io"
	"os"
)

// We want to write a function that greets someone, like in the hello-world section,
// but this time we are going to be testing *the actual printing*

// Here is what the function could look like
// func Greet(name string) {
// 	fmt.Printf("Hello, %s", name)
// }

// but this is hard to test as `fmt.Printf` prints to `stdout` which is hard to capture using the testing framework.
// What we need to do is `inject` (fancy word for pass-in) the dependency of printing.

// The function doesn't need to care *where* or *how* the printing happens, so we should accept an `interface` instead of a concrete type.
// If we do that then we can change the implementation to print something we control so that we can test it.
// in "real life" youwould inject in something that writes to stdout.

// in the Go library, writing to stdout uses the function `os.Stdout` which implements `io.Writer`, the Writer interface looks like this:
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// Writer.Write expects an array of bytes, so we can try using this in out tests
// `fmt.Printf` defaults to stdout. So use `fmt.Fprintf` which takes a `Writer` to send the string to.
// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// if we try to call Greet to print to Stdout we get an error:
// cannot use os.Stdout (variable of type *os.File) as *bytes.Buffer value in argument to Greet,
func main() {
	Greet(os.Stdout, "Adam")
}

// Therefore we should change the Greet function to accept `io.Writer` instead of `bytes.Buffer`
// as both `os.Stdout` and `bytes,Buffer` implement it.
// and now the main() function will work: `go run di.go`
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Summary:
// Injecting a dependency allows us to control *where* the data was written two which allows us to test the function easily.
// It allows us to separate concerns - decoupling *where the data goes* from *how to generate it*
// Also our code can be re-used in different contexts. The first "new" context is inside tests, but further on someone could
// try something new and inject their own dependencies to the Greet function.

// It is worth studying the `io.Writer` interface in the standard library documentation - it can be used in many different contexts such as
// in writing repsonses in the http package.
