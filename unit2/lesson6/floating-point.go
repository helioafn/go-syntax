// Lesson 6.1: Declaring floating-point variables

package main

import "fmt"

func main() {
	// var days float64 = 365.2425
	// var days = 365.2425

	// Equivalent to both above, here the Go compiler will infer float64 to days
	// days := 365.2425
	var answer float64 = 42
	fmt.Println(answer)
}
