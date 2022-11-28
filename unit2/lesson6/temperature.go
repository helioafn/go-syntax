// Listing 6.7: Multiplication first

package main

import "fmt"

func main() {
	celsius := 21.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32
	fmt.Println(fahrenheit, "ºF") // 69.8ºF

	// Quick check 6.5: What is the best way to avoid rounding errors?
	// A: Don't use a floating-point
}
