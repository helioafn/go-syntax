// Listing 6.3: Formatted print for floating-point
package main

import "fmt"

func main() {
	// %f formatting verb to specify the number of digits
	// %2.3f -> 2 means the width(minimal number of chars to display, including the decimal point)
	// and .3 the precision after the decimal point
	third := 123.45678
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%2.3f\n", third)
	fmt.Printf("%.4f\n", third)

	// Quick check 6.4.2
	// A: width is 9, and precision is 4
}
