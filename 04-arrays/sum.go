package main

// Arrays have a fixed capacity which you define when you declare the variable.
// they can be initialised in two ways:
// `numbers := [5]int{1, 2, 3, 4, 5}` or `numbers := [...]int{1, 2, 3, 4, 5}`
// For arrays, the size is encoded in it's type, so `[4]int` is not compatible with `[5]int`
// Go has Slices which do not encode the size of the collection and instead can have any size

// Step 1. Minimum code for test to run and check the failing test output:
// func Sum(numbers [5]int) int {
// 	return 0
// }

// Step 2. Enough code to make the test pass
// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

// Step 3. Refactor
// Notice we have changed to using a Slice now `[]int`,
// so that we can have an arbitrary size of the numbers array input
func Sum(numbers []int) int {
	sum := 0
	// `range` iterator returns an `index` and `value`
	// we don't need the index so ignore it using `_` (blank identifier)
	for _, number := range numbers {
		sum += number
	}
	return sum
}
