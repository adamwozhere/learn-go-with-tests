package numerals

import "strings"

// func ConvertToRoman(arabic int) string {

// 	// a `Builder` is used to efficiently build a string using `Write` methods,
// 	// it minimizes memory copying - It's not much larger than "manual" append-ing,
// 	// so we might as well use the faster approach.
// 	var result strings.Builder

// 	// try just adding I for now
// 	// for i := 0; i < arabic; i++ {
// 	// 	result.WriteString("I")
// 	// }

// 	// we need to be able handling 4 as IV not IIII etc. (can't repeat more than 3 times)
// 	// for i := arabic; i > 0; i-- {
// 	// 	if i == 5 {
// 	// 		result.WriteString("V")
// 	// 		break
// 	// 	}
// 	// 	if i == 4 {
// 	// 		result.WriteString("IV")
// 	// 		break
// 	// 	}
// 	// 	result.WriteString("I")
// 	// }

// 	// we're repeating the same logic so refactor:
// 	// given the understanding from the code driven from the tests for 1 to 5 so far:
// 	// we can see that to build a roman numeral we need to subtract from `arabic` as we apply symbols.
// 	// So the foor loop no longer relies on `i` and instead we keep building the string until we have
// 	// subtracted enough symbols away from `arabic`.
// 	for arabic > 0 {
// 		switch {
// 		case arabic > 4:
// 			result.WriteString("V")
// 			arabic -= 5
// 		case arabic > 3:
// 			result.WriteString("IV")
// 			arabic -= 4
// 		default:
// 			result.WriteString("I")
// 			arabic--
// 		}
// 	}
// 	// its obvious we need to refactor here - but add more tests till we can see the pattern.
// 	// According to some, in OO programming you should view switch statements with a bit of suspicion -
// 	// usually you are capturing a concept or data inside some imperitive code, when in fact it
// 	// could be captured in a class structure instead.
// 	// Go isn't strictly OO but that doesn't mean we ignore the lessons OO offers entirely.

// 	return result.String()
// }

// refactoring in an "OO" style:
// the switch statement is describing some truths about roman numerals along with behaviour,
// lets decouple the data from the behaviour:

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var alLRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	// extended set of symbols for full tests above:
	// the algorithm didn't change, only extra data added.
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {

	var result strings.Builder

	for _, numeral := range alLRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

// this feels better - we've just declared some rules around the numerals as data,
// rather than hidden in an algorithm, and we can see that we just go through the arabic number
// adding symbols if they fit.

// now we need to add a function to convert from a roman numeral to an int.

func ConvertToArabic(roman string) uint16 {
	// in the same way, we need to iterate through the input and build the total
	// total := 0
	// for range roman {
	// 	total++
	// }

	// return total

	// instead we can do the same algorithm but in reverse:
	var arabic uint16 = 0

	for _, numeral := range alLRomanNumerals {
		// go through each numeral in allRomanNumerals
		// then essentially while loop, checking if the symbol is at the start of the string,
		// if so, increment the arabic integeger and remove the symbol from the input string
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}

// SUMMARY:
// Here we can see that TDD really does help when doing something complicated,
// we just have to remember that to start with something simple and take small steps -
// don't make any large leaps or do huge refactorings.
// This method works for big tasks - we just have to know *how* to split work up
// and iterate accordingly.
// Property tests are built into the standard testing library and give confidence,
// especially when noticing and enforcing "domain rules" (although we only tested the conversion-to-and-back-again rule here!)
