// 3.1 True of false
package main

import (
	"fmt"
	"strings"
)

func main() {
	//  Listing 3.1
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)

	// Quick check 3.1.1
	var wearShades = true
	fmt.Println("Have you wore you sunglasses?", wearShades)

	// Quick check 3.2.2
	var isSignRead = strings.Contains(command, "read")
	fmt.Println("Did you read the sign at the entrance?", isSignRead)

}
