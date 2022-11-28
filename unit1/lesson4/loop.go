// Listing 4.2: A condensed countdown

package main

import "fmt"

func main() {
	// Short declaration: it provides an alternative syntax for the var keyword
	//var count = 10  // equivalent to the line below
	count := 0

	// More than just less characters, short declaration can go places where var can't
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println(count) // Still in the scope

	// Without short declaration, count variable must be declared outside of the loop,
	// memaning it remains in scope after the loop ends.

}
