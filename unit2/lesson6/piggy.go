/*
	Experiment:
		Save some money to buy a gift for your friend. Write a program that randomly places
		nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it con-
		tains at least $20.00. Display the running balance of the piggy bank after each deposit,
		formatting it with an appropriate width and precision.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	savings := 0.0

	for savings < 20.0 {
		switch coins := rand.Intn(3); coins {
		case 0:
			savings += 0.05
		case 1:
			savings += 0.10
		case 2:
			savings += 0.25
		}
		fmt.Printf("Balance: $%.2f\n", savings)
	}
}
