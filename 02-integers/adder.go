package integers

// first write the minimum code to satisfy the compiler,
// allowing the test to run and check the failing test output e.g.

// func Add(x, y int) int {
// 	return 0
// }

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

// we are just returning an int here, not a `named return value`.
// A named return value should generally be used when the meaning of the result
// isn't clear from context - here it's pretty clear that Add will add the parameters.

// a named return value should appear in the documentation generated and in the text editor,
// the comment above a function will also appear in the documentation with `godoc`.

// for more info: https://go.dev/wiki/CodeReviewComments#named-result-parameters
