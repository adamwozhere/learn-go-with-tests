package main

import "fmt" // the format package handles formatting like printing strings, and also for fotmatting code!

// separate "domain" code from outside world (side-effects)
// the fmt.Println is a side effect (for printing to stdout)
// and the string we send is our domain

const englishHelloPrefix = "Hello, "

// function with return type of string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
