package main

import (
	"testing"
)

// func TestWallet(t *testing.T) {

// 	wallet := Wallet{}
// 	wallet.Deposit(10)

// 	got := wallet.Balance()
// 	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

// 	want := 10

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// refactoring test using Bitcoin type
func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		// notice with the Wallet struct you use curly braces {}
		wallet := Wallet{}

		// whereas Bitcoin type you use normal braces (),
		// maybe because it's underlying type is primitive as it's an int?
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		// making sure that we check if Withdraw is successful, that NO error is returned
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		// we need to add a return type to Withdraw for this to work
		err := wallet.Withdraw(Bitcoin(100))

		// check for both error, and balance (should be the same as startingBalance)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

// This time helpers are moved out of the main test so a developer can see the test assertions first rather than the helpers

// note that `testing.TB` is used as "B" interface is needed for `t.Helper`
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// assertError := func(t testing.TB, err error) {
// 	t.Helper()

// 	// `nil` is like `null` from other languages.
// 	// Errors can be `nil` because the return type of `Withdraw` will be `error`, which is an interface.
// 	// If you see a function that takes arguments or returns values that are interfaces, they can be nillable.
// 	// like `null`, if you try to access a value that is `nil`, it will throw a runtime panic. Make sure to check for nils.
// 	if err == nil {
// 		t.Error("wanted an error but didn't get one")
// 	}
// }

// refactor the assertError function so it can iterate / assert the kind of error message rather than just the existance of an error.
func assertError(t testing.TB, got error, want error) {
	t.Helper()

	// check for the presence of an error first,
	// if there is no error, then `t.Fatal` will stop the test,
	// preventing a panic when trying `got.Error()`
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	// `error.Error()` method converts the error to a string, if we wanted to compare it to a message string,
	// but as we have the error as a global package variable we can just compare them both directly
	if got != want {
		// `%q` a double-quoted string safely escaped with Go syntax, used as our error message is quoted?
		t.Errorf("got %q wanted %q", got, want)
	}
}

// The Go compiler helps a lot but sometimes there are things that can be missed:
// there is one scenario that is not tested. To find it install "errcheck":
// `go install github.com/kisielk/errcheck@latest` and run `errcheck .` in the terminal.
// it will print e.g. `wallet_test.go:38:18:   wallet.Withdraw(Bitcoin(10))`.
// This is because we have not checked that if Withdraw is successful, that an error is NOT returned!
// adding a function for this

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		// this is also fatal so it stops the tests
		t.Fatal("got an error but didn't want one")
	}
}
