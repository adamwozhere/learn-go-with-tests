package main

import "fmt" // the format package handles formatting like printing strings, and also for fotmatting code!

// separate "domain" code from outside world (side-effects)
// the fmt.Println is a side effect (for printing to stdout)
// and the string we send is our domain

const spanish = "Spanish"
const french = "French"
const german = "German"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "

// function with return type of string
func Hello(name, language string) string {
	if name == "" {
		name = "World"

	}
	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
