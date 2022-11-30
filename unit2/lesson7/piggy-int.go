/*
	Experiment
	Write a new piggy bank program that uses integers to track the number of cents rather
	than dollars. Randomly place nickels (5¢), dimes (10¢), and quarters (25¢) into an empty
	piggy bank until it contains at least $20.
	Display the running balance of the piggy bank after each deposit in dollars (for exam-
	ple, $1.05).
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0.0
	cents := 0

	for piggyBank < 20.0 {
		switch coins := rand.Intn(3); coins {
		case 0:
			cents += 5
		case 1:
			cents += 10
		case 2:
			cents += 25
		}

		// 100 cents == 1 dollar
		if cents >= 100 {
			piggyBank += 1.0
			cents -= 100                        // to get remaining cents
			piggyBank += float64(cents) / 100.0 // adding the remaining cents
			fmt.Printf("$%.2f\n", piggyBank)
		}
	}
}
