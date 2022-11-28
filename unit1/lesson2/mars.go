// My weight loss program
package main

import (
	"fmt"
)

// Entry point of the programs
func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(85.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(25 * 365 / 687)
	fmt.Println(" years old.")

	// The above can be achieved like this, passing a list of arguments to Println (text, number, math expressions)
	fmt.Println("My weight on the surface of Mars is", 85.0*0.3783, "lbs, and i would be",
		25*365/687, "years old")
}
