package iteration

// Repeat takes a character and returns it `count` times
func Repeat(character string, count int) string {
	// before we have used `:=` to declare an initialise variables,
	// it is a shorthand for doing both steps.
	// here we are declaring a string variable only,
	// so are using the explicit version with the `var` keyword.
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

// alternatively we could use built in library functions.
// for example by importing "strings", we could simply do `return strings.Repeat(character, count)`
