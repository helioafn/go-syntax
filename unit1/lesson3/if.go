package main

import "fmt"

func main() {
	// Listing 3.3
	var command = "go east"

	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life")
	} else {
		fmt.Println("Didn't quite get that")
	}

	// Quick check 3.3
	var currentRoom = "mountain"

	if currentRoom == "entrance" {
		fmt.Println("At the entrance has this gorgeous gate and you can hear clearly the sound of the wind")
	} else if currentRoom == "cave" {
		fmt.Println("Passing by the entrance and take the road to the left, we get to this exquisite cave",
			"feels like something is around here.")
	} else if currentRoom == "mountain" {
		fmt.Println("Finally you arrived at the highest spot from the trip,",
			"a amazing mountain with a incredible view")
	}
}
