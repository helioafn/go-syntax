// Listing 6.1: 64-bit vs 32-bit floating-point
package main

import (
	"fmt"
	"math"
)

func main() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64) // 3.141592653589793
	fmt.Println(pi32) // 3.1415927

	// ! Functions in the math package operate on float64 types, so prefer using it, unless you have
	// a good reason to do otherwise.

	// Quick check 6.2: How many bytes of memory does a single precision float32 use?
	// A: it uses 4 bytes = 32 bits of memory.
}
