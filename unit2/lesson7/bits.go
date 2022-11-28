// Listing 7.3 - Display the bits

package main

import "fmt"

func main() {
	/*
		To understand why integers wrap, take a look at the bits.
		The %b format verb will show the bits for an integer value.
		Like other format verbs, you can pad and add a minimum length
	*/

	var green uint8 = 3
	fmt.Printf("%08b\n", green)

	green++
	fmt.Printf("%08b\n", green)
}
