// Listing 6.6: Division-first
package main

import "fmt"

func main() {
	celsius := 21.0
	fmt.Println((celsius/5.0*9.0)+32, "º F")
	fmt.Println((9.0/5.0*celsius)+32, "º F")
	// Both lines above prints: 69.80000000000001 º F
}
