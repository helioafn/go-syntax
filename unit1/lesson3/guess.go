/*
	Experiment: guess.go
	Write a guess-the-number program. Make the computer pick random numbers
	between 1-110 until it guess your number, which you declare at the top of
	the program. Display each guess and wheter it was too big or too small.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const myChoice = 77

	for {
		var computerChoice = rand.Intn(101)
		if computerChoice == myChoice {
			fmt.Println("You guessed my choice! It was:", computerChoice)
			break
		} else if computerChoice < myChoice {
			fmt.Println("Your guess:", computerChoice, ". It was too small")
		} else if computerChoice > myChoice {
			fmt.Println("You guess:", computerChoice, ". It was too big")
		}
	}
}
