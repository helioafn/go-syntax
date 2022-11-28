// Listing 7.4 - The bits when integers wrap
package main

import "fmt"

func main() {
	var blue uint8 = 255
	fmt.Printf("%08b\n", blue) // 11111111
	blue++
	fmt.Printf("%08b\n", blue) // 00000000

	// Wrapping may be what you want in some situations, but not always.
	// The simplest way to avoid wrapping is to use an integer type large enough to hold the values
	// you expect to store

}
