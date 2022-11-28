// Listing 6.2: Declaring a variable without a variable

/*
	The zero value: In Go, each type has a default value. It applies when you declare a variable
	but don't initialize it with a value.
*/

package main

import "fmt"

func main() {
	var price float64 // to the computer is identical to the following -> price := 0.0
	fmt.Println(price)

	// Quick check 6.3: What is the zero value for a float32?
	// A: 0 (zero)
}
