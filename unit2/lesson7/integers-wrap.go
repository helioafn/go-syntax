// Listing 7.2 - Integers wrap around
package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Integers are free of the rounding errors that make floating-point inaccurate,
		but integers have a different problem: a limited range.
		In Go, when that range are exceeded integers wrap around.
	*/

	var red uint8 = 255
	red++
	fmt.Println(red) // 0

	var number int8 = 127
	number++
	fmt.Println(number) // -128

	// Quick check 7.5.1
	red = 255
	red += 20
	fmt.Println(red) // 19

	// Quick check 7.5.2
	number = -128
	number -= 33
	fmt.Println(number) // 95

	// Quick check 7.5.3
	var num int16 = math.MaxInt16
	var num2 uint16 = math.MaxUint16
	var num3 int32 = math.MaxInt32
	var num4 int64 = math.MaxInt64
	fmt.Println(num + 1)  // -32768
	fmt.Println(num2 + 1) // 0
	fmt.Println(num3 + 1) // -2147483648
	fmt.Println(num4 + 1) // -9223372036854775808
}
