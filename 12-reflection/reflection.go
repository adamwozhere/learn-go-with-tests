package main

import "reflect"

// Challenge: write a function `walk(x interface{}, fn func(string))` which takes a struct `x`
// and calls `fn` for all strings fields found inside.

// to do this we need to use "reflection" - it is the ability for a program to examine it's own structure,
// particularly through types; it's a form of metaprogramming. It's also a great source of confusion.

// `interface{}` is basically "any type" - go actually has an alias `any` for it.
// This allows the walk function to accept any value for `x`.

// 1. basic function to get the test to compile
// func walk(x interface{}, fn func(input string)) {
// 	fn("test")
// }

// 2. enough code to make it pass
// func walk(x interface{}, fn func(input string)) {
// 	// we need to use reflection to try and look at `x`'s properties
// 	// the reflect package allows us to do this - `ValueOf` returns the Value of a given variable,
// 	// which we can then inspect the `Field`'s.
// 	// This passes the inital test but it is not safe as it checks only the first field and assumes it's a string.
// 	val := reflect.ValueOf(x)
// 	field := val.Field(0)
// 	fn(field.String())
// }

// // 3. refactor
// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		// we need to check that the field is a string
// 		// not sure if we always have to use reflect.[Type] in comparisons,
// 		// or if in other situations types can just be compared e.g. "string" == string.
// 		if field.Kind() == reflect.String {
// 			fn(field.String())
// 		}

// 		// we also need to make sure that nested fields work
// 		if field.Kind() == reflect.Struct {
// 			// call walk again on this  inner struct, Note: we need to use `.Interface()`
// 			walk(field.Interface(), fn)
// 		}
// 	}
// }

// // 4. Refactor again
// // generally when doing a comparison on the same value more than once, a `switch` would be more readable and easy to extend.
// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	// make sure that we can use pointers - we can't use NumField on a pointer Value,
// 	// so extract the underlying value first with `Elem()`
// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		// note that cases automatically break in Go, if you want them to fall through, you can use the `fallthrough` keyword
// 		switch field.Kind() {
// 		case reflect.String:
// 			fn(field.String())
// 		case reflect.Struct:
// 			walk(field.Interface(), fn)
// 		}
// 	}
// }

// 5. Refactoring further - more code but separates the concerns better,
// also we want to `walk` on either each field in a struct, or each "thing" in a slice -
// so lets write the code to better reflect that and switch on the type first.
// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	switch val.Kind() {
// 	case reflect.Struct:
// 		for i := 0; i < val.NumField(); i++ {
// 			walk(val.Field(i).Interface(), fn)
// 		}
// 	case reflect.Slice:
// 		for i := 0; i < val.NumField(); i++ {
// 			walk(val.Index(i).Interface(), fn)
// 		}
// 	case reflect.String:
// 		fn(val.String())
// 	}
// }

// 6. Refactoring further - there is repetition in the loops in each case
// as it's getting the number of values and accessing them which changes depending on the type,
// we can use the switch to get them appropriately and assign to variables,
// then loop over them after.
// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	numberOfValues := 0

// 	// this function expression is a little bit confusing to me;
// 	// I thought you would have to call it like `getField(x)`,
// 	// but actually it looks like we are just telling getField takes a `func(int)` ?
// 	// so then later we can assign e.g. `val.Index` to it, and when we finally call `getField(i)` -
// 	// it is actually `val.Index(i)`.
// 	var getField func(int) reflect.Value

// 	switch val.Kind() {
// 	case reflect.String:
// 		fn(val.String())
// 	case reflect.Struct:
// 		numberOfValues = val.NumField()
// 		getField = val.Field
// 		// we can use the same code to handle slices AND arrays
// 	case reflect.Slice, reflect.Array:
// 		numberOfValues = val.Len()
// 		getField = val.Index
// 		// maps are similar to structs, but the keys are unknown at compile time
// 	case reflect.Map:
// 		// You can't get values out of a map by index, only by key, so it breaks our abstraction
// 		for _, key := range val.MapKeys() {
// 			walk(val.MapIndex(key).Interface(), fn)
// 		}
// 	}

// 	for i := 0; i < numberOfValues; i++ {
// 		walk(getField(i).Interface(), fn)
// 	}
// }

// 7. Refactoring to use maps - the previous version felt like a nice abstraction, but now with maps it doesn't feel goo.
// -- it's ok to make mistakes and change the abstraction while doing TDD - it gives us freedom to try things out.
// lets take things back to how they were before with looping within the switch cases
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// DRY the code a bit with a function to call walk inside the switch cases so that they only have to extract
	// the `reflect.Value`s from `val`.
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
		// add support for channels
	case reflect.Chan:
		// NOTE: unsure of this for loop syntax so need to look into it.
		// I think it's essentially a while loop, that gets the `v` when it's recieved from a channel and runs walkValue on it,
		// - unsure if the loop closes when it runs out of `v` values and triggers break.
		// or if the `ok` value is received when finished which triggers break.
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
		// functions
	case reflect.Func:
		// assuming that you "call" the function with nil as it takes no arguments, then iterate on the returned values
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

// this function now just gets the value correctly - whether it's a pointer or not
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

// Summary:
// Reflection allows a program to examine it's own structure, particularly through types.
// Is a form of metaprogramming.
// I'll need to look into this more as it's quite a confusing topic.
