package numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

// This chapter shows you how to tackle the Roman Numeral Kata with TDD
// We will write a function which convers an Arabic number (0-9) to a Roman Numeral.

// Remember that a key skill is to try and identify "thin vertical slices" of *useful* functionality,
// and then iterating - the TDD workflow helps facilitate this:
// so rather than 1984, lets start with 1.

// func TestRomanNumerals(t *testing.T) {
// 	// got := ConvertToRoman(1)
// 	// want := "I"

// 	// if got != want {
// 	// 	t.Errorf("got %q, want %q", got, want)
// 	// }

// 	// We have a lot of tests to wrie - whenever we have repetition, if it feels like "given input X, we expect Y",
// 	// we should probably use table based tests.

// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
// 		{"5 gets converted to V", 5, "V"},
// 		{"9 gets converted to IX", 9, "IX"},
// 		{"10 gets converted to X", 10, "X"},
// 		{"14 gets converted to XIV", 14, "XIV"},
// 		{"18 gets converted to XVIII", 18, "XVIII"},
// 		{"20 gets converted to XX", 20, "XX"},
// 		{"39 gets converted to XXXIX", 39, "XXXIX"},
// 		{"40 gets converted to XL", 40, "XL"},
// 		{"47 gets converted to XLVII", 47, "XLVII"},
// 		{"49 gets converted to XLIX", 49, "XLIX"},
// 		{"50 gets converted to L", 50, "L"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// it looks like we can now refactor the test a bit better, and use a data structure like in the "OO" version of the function.

// move the cases outside the test as package variable so we can use them for the roman to arabic tests also
var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		// we don't really need a "description" as the arabic and roman explain the test.
		// `Sprintf` formats according to specifier and returns the string - so this becomes the test description
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

// now we add tests for converting roman numerals to arabic
func TestConvertingToArabic(t *testing.T) {
	// notice - just taking a slice of ONE case for now to test if it actually works `cases[:1]`
	// then 4 cases etc. increasing as we get more tests to pass.
	for _, test := range cases[:4] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

// Property based tests
// we notice that there are a few rules in the domain of Roman Numerals:
// can't have more than 3 consecutive symbols,
// only I (1), X, (10), and C (100) can be "subtractors"
// taking the result of `convertToRoman(N)` and passing it to `convertToArabic` should return N.
// The tests we have done so far can be described as "example" based tests where we provide *examples*
// for the tooling to verify.
// We can take these rules that we know about our domain and use property based tests: these throw random
// data at out code and verify that the domain rules you describe are true.
// you could think that htese tests are mainly about random data but this is not true - the challenge about
// property based tests is having a *good* understanding of your domain so you can write these properties.

// func TestPropertiesOfConversion(t *testing.T) {
// 	// check that a number converted to a numeral, then back to a number is the same.
// 	assertion := func(arabic int) bool {
// 		// the function was hanging, but if we put this code, we can see that negative numbers,
// 		// and numbers too large to be expressed as roman numerals (according to our "no more than consecutive 3 symbols" rule)
// 		// are passed into the function.
// 		if arabic < 0 || arabic > 3999 {
// 			log.Println(arabic)
// 			return true
// 		}
// 		roman := ConvertToRoman(arabic)
// 		fromRoman := ConvertToArabic(roman)
// 		return fromRoman == arabic
// 	}

// 	// `Check` runs a function with random data, it must be a function that returns a bool.
// 	if err := quick.Check(assertion, nil); err != nil {
// 		t.Error("failed checks", err)
// 	}
// }

// we could use a `uint16` which is an unsigned integer, and then just return true if the number is > 3999.
// we must also remember to change our actual functions to use uint16 now too instead of an int!

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		// if the number is out of upper bounds, just effectively skip it.
		if arabic > 3999 {
			return true
		}

		// we can print what actual values are being tested with log
		// and running command `go test -v` to print the addional output (`-v` flag)
		t.Log("testing", arabic)

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	// the `Check` function defaults to 100 test counts, but we can define our own number.
	// in go we can "do something" and check it in an if statement at the same time - a one liner if, else statement:
	// `if` keyword followed by an initialisation statement, a `;` semicolon, followed by the condition.
	// e.g. `if err := doStuff(); err != nil { ...handle error here }`.
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

// There are further tests we could do, make sure we check each of the other domain properties we described
// (although we know now that with a combination of the original "example" tests and the "property" tests that our function works)
// we could also limit the numeric imput of out function so that it doesn't exceed 3999:
// either by returning an error, or create a new type that cannot represent > 3999 -- although I think you would still
// have to have an some sort of constructor function that checks for the range and returns an error?
