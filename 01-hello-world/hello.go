package main

import "fmt" // the format package handles formatting like printing strings, and also for fotmatting code!

// separate "domain" code from outside world (side-effects)
// the fmt.Println is a side effect (for printing to stdout)
// and the string we send is our domain

const (
	spanish = "Spanish"
	french  = "French"
	german  = "German"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	germanHelloPrefix  = "Hallo, "
)

// function with return type of string
// the function name begins with a capital letter which means it is a public function
// to make a function private, use a lowercase starting letter
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// function signature has a `named return value` (prefix string)
// this creates a variable called `prefix` in the function,
// it will be assigned a "zero" value depending on type,
// and can be returned simply with `return` rather than `return prefix`
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
