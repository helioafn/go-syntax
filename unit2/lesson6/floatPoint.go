// Comparing floating-point numbers

package main

import (
	"fmt"
	"math"
)

func main() {
	piggyBank := 0.0
	for i := 1; i <= 11; i++ {
		piggyBank += 0.1
	}

	fmt.Println(piggyBank)

	// Taking the absolute difference between two numbers
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001) // true
}
