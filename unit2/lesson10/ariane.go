// Listing 10.2 - Ariane type conversion
package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32768
	var h = int16(bh)
	fmt.Println(h)

	// "Handling" values outside an range
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		// do something
	}
	// These min/max constants are untyped, allowing the comparison of bh, a floating point value, to MaxInt16.

	// Quick check 10.3
	v := 1337

	if v >= 0 && v <= math.MaxUint8 {
		vConverted := uint8(v)
		fmt.Printf("%v is in %[1]T range\n", vConverted)
	} else {
		fmt.Printf("%v it's outside of uint8 range\n", v)
	}
}
