/*
	Experiment:
	Write a program that converts strings to Booleans:
		* The strings “true”, “yes”, or “1” are true;
		* The strings “false”, “no”, or “0” are false;
		* Display an error message for any other values.
*/

package main

import "fmt"

func main() {
	var converted bool
	str := "1"

	switch str {
	case "true", "yes", "1":
		converted = true
	case "false", "no", "0":
		converted = false
	default:
		fmt.Println("Invalid value!")
	}

	fmt.Println(converted)
}
