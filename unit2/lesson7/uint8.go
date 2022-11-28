// 7.2 - The uint8 type for 8-bit colors

package main

import "fmt"

func main() {
	// Lines 9, 10 are equivalent
	var red, green, blue uint8 = 0, 141, 213
	red, green, blue = 0x00, 0x8d, 0xd5        // Go requires a 0x prefix for hexadecimal
	fmt.Printf("%x %x %x\n", red, green, blue) // %x (or %X for uppercase) format verb display numbers in hexadecimal

	// Outputing color css-like.
	// As with %v and %f verbs, you can specify a minimum number of digits and zero padding with %02x(%02X)
	fmt.Printf("color: #%02X%02X%02X;\n", red, green, blue)
}
