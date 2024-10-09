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

// Sum returns all the numbers in a slice added together.
func Sum(numbers []int) int {
	sum := 0
	// `range` iterator returns an `index` and `value`
	// we don't need the index so ignore it using `_` (blank identifier)
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// input any number of integer arrays, and return an integer array with the sum of each array
// Step 1. Minimum code
// func SumAll(numbersToSum ...[]int) []int {
// 	return nil
// }

// Step 2. Make it pass
// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)

// 	// use `make` to create a slice with initial length of `lengthOfNumbers`
// 	// The length is the number of elements it holds,
// 	// while the capacity is the number of elements it can hold in the underylying array `cap(mySlice)`,
// 	// e.g. `make([]int, 0, 5)` creates a slice with length 0 and capacity 5.

// 	sums := make([]int, lengthOfNumbers)

// 	// loop through the range, putting the sum into the index of the sums array

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

// Step 3. Refactor
// In the previous example we used the index, which could give a runtime error if index is out of range.
// Now we use the `append` function which takes a slice and a new value,
// then returns a new slice with all the items in it.
// In this implementation we don't worry about capacity, we start with an empty slice `sums` and append to it.
// func SumAll(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		sums = append(sums, Sum(numbers))
// 	}

// 	return sums
// }

// New requirement is to change SumAll to SumAllTails, where it will calculate the totals of the "tails"
// of each slice. The tail of a collection is all items in the collection except the first one (the "head")

// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		// slices can be sliced with syntax `slice[low:high]`.
// 		// ommitting the value on one side means it will take everything on that side.
// 		// here we are saying "take from 1 to the end" (leaves off index 0)
// 		tail := numbers[1:]
// 		sums = append(sums, Sum(tail))
// 	}

// 	return sums
// }

// rewrite the function so that it passes if an empty slice is passed
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
