// Listing 7.1 - Inspect a variable's type

package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year) // int (on this environment case int64)

	// Instead of repeating the variable twice, you can tell Printf to use the first argument [1]
	// for the second format verb
	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days) // float64

	// Quick check 7.3
	whoAmI := "Hélio"
	fmt.Printf("Type %T for %[1]v\n", whoAmI) // string
	fmt.Printf("Type %T for %[1]v\n", true)   // bool
}
