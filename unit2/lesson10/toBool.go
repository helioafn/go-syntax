// Listing 10.6 - Converting string to a Boolean
package main

import "fmt"

func main() {
	yesNo := "no"
	// 	The inverse conversion (string to Boolean) requires less code because you can assign the result of a condition directly
	// 	to a variable.
	launch := (yesNo == "yes")
	fmt.Println("Ready for launch:", launch) // Ready for launch: false

	/* The Go compiler will report an error if you attempt to conver a Boolean with string(false), int(false) or similar,
	and likewise for bool(1) or bool("yes"):
		someValue := string(false) cannot convert false (untyped bool constant) to type string
		someValue1 := int(false) cannot convert false (untyped bool constant) to type int
		someValue2 := bool(1) cannot convert 1 (untyped int constant) to type bool
		someValue3 := bool("yes") cannot convert "yes" (untyped string constant) to type bool
	*/

	// Boolean in Go don't have a numeric equivalent

	// Quick check 10.5
	boolValue := true
	var oneZero int

	if boolValue {
		oneZero = 1
	} else {
		oneZero = 0
	}
}
