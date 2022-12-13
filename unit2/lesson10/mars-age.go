// Listing 10.1 - Mars age
package main

import "fmt"

func main() {
	age := 25
	marsAge := float64(age) // converting to float64

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.") // I am 13.291211790393014 years old on Mars.

	/*
		Type conversion are required between unsigned and signed integer types, and between types of
		different sizes. It's alwas safe to conver to a type with a larger range, such as int8 to an int32.
		Other integer conversions come with some risks. A uint32 could contain a value of 4 billion, but an
		int32 only supports numbers to just over 2 billion. Likewise, an int may contain a negative number, but
		a uint can't.
	*/

	/*
		Quick check 10.2
		A1: uint8(red)
		A2: invalid operation: type mismatch (int and float64)
	*/
}
