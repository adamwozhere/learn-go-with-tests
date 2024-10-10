package main

import "math"

// func Perimeter(width float64, height float64) float64 {
// 	return 2 * (width + height)
// }

// func Area(width, height float64) float64 {
// 	return width * height
// }

// The above functions work, but it wouldn't work if someone passed in values for a triangle etc.
// We can create a simple type using a `struct`.
// A struct is just a named collection of fields where you can store data.

type Rectangle struct {
	Width  float64
	Height float64
}

// Now the struct can be used as the argument for the functions

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

// some programming languges will allow you to declare the Area function twice,
// with one taking Rectangle as an argument and another taking Circle, but you cannot do this in Go.
// You can either have the two functions in two different packages, but that is overkill here.
// We can define methods on our types instead.
// A method is a function with a receiver. A method declaration binds an identifier, the method name,
// to a method, and associates the method with the receiver's base type.
// Methods are like functions but they invoked on a particular type
// (you can only call methods on "things" - e.g. `Errorf` is a method on an instance of `testing.T`)

// Here are the methods on the types
// notice the syntax for a method is similar to a function but with a receiver:
// `func (receiverName ReceiverType) MethodName(args) return type {}`
// so here we can see that the method Area() is bound to the type Rectangle,
// in other languages, the methods are written inside the class / object / struct,
// and the receiver is accessed with the `this` keyword rather than e.g. `r`.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Interfaces allow you to make functions that can be used with different types,
// and create highly-decoupled code whilst still maintaining type safety

type Shape interface {
	Area() float64
}

// in other programming languages you would normally have to write:
// `My type Foo implements interface Bar`.
// But in this case, both `Rectangle` and `Circle` have a method called `Area` that returns a `float64`,
// so they both satisfy the `Shape` interface.
// in Go, interface resolution is implicit - if the type you pass in matches what the interface
// is asking for, it will compile.

// It's now easy to add a new Triangle Shape and test it.

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// using structs and interfaces in this way to use to declare functions that can be used by different types
// is known as "parametric polymophism"
