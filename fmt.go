// Presenting fmt.Printf()
package main

import "fmt"

func main() {
	// Unlike Print and Println, the first argument to Printf is always text
	// %v is a format verb, it substitutes for the value provided by the second argument
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 85.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 25*365/687)

	// If multiple format verbs are specified, Printf will substitute multiple values in order
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)

	// You can align text by specifying width as part of the format verb (ex.: %4v), to pad a value to
	// a width of 4 characters. * Positive number pads with spaces to the left, negative numbers pads
	// with spaces to the right
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
