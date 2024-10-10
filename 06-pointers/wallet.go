package main

import (
	"errors"
	"fmt"
)

// type Wallet struct {
// 	// this variable is a lowercase-starting, so it is private outside the package
// 	balance int
// }

// Use methods to manipulate the wallet ballance, (similar to getter/setter)
// func (w Wallet) Deposit(amount int) {
// 	w.balance += amount
// }

// func (w Wallet) Balance() int {
// 	// `%p` prints pointer value in hex
// 	// `&` before a variable gets the memory address of that variable
// 	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
// 	return w.balance
// }

// The above does not work as in Go, when you call a function or method,
// the arguments are COPIED - running the test you can see the two different memory addresses.

// to fix this we use a pointer, `*` before a variable.
// here we use the receiver type `*Wallet`,
// which can be read as "a pointer to a wallet".

// func (w *Wallet) Deposit(amount int) {
// 	w.balance += amount
// }

// func (w *Wallet) Balance() int {
// 	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
// 	return w.balance
// }

// technically you don't need to change `Balance` to use a pointer receiver as taking a copy of the balance is ok.
// However, by convention you should keep your method receiver types the same for consistency.

// pointers to structs have their own name: "struct pointers" and are automatically dereferenced
// (I don't understand this so need to look into dereferencing)

// Refactoring with a Bitcoin type

// types can be created from existing ones
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// to make `Bitcoin` you just use the syntax `Bitcoin(999)`,
// By doing this we're making a new type and can declare methods on them. This can be useful when you
// want to add some domain specific functionality on top of existing types.

// This Stringer interface is already defined in the `fmt` package
// It allows you to define how your type is printed when used with the `%s` format string in prints
type Stringer interface {
	String() string
}

// add a method to our Bitcoin type
func (b Bitcoin) String() string {
	// `Sprintf` formats according to a format specifier
	// `%d` specifies base 10 format
	return fmt.Sprintf("%d BTC", b)
}

// errors in Go are values, unlike in JavaScript etc where we would have to try / catch and throw an error
// so we can return an `error` from this function, which is `nil` if there's no error.

// instead of `return errors.New("cannot withdraw, insufficient funds")`, lets create a error variable,
// this can then also be used in the test, and means we have one source of truth

// the var keyword allows us to define values global to the package (so we can use it in tests)
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		// `errors.New` creates a a new error with message
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Summary:
// Go copies valies when you pass them to function/methods, so if you need to mutate state then use a pointer to that state
// pointers can be nil, you must check if it's nil otherwise it might cause a runtime exception.
// Errors are the way to signify failure when calling a function/method. They are values. Check for errors and handle them.
