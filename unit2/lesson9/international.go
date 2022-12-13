/*
	Experiment:
	Cipher the Spanish message “Hola Estación Espacial Internacional” with ROT13. Modify
	listing 9.7 to use the range keyword. Now when you use ROT13 on Spanish text, characters
	 with accents are preserved.
*/

package main

import "fmt"

func main() {
	messageToBeCiphered := "Hola Estación Espacial Internacional"
	// messageToBeDeciphered := "Hbyn Efgnpvóa Efcnpvny Iagreanpvbany"
	for _, char := range messageToBeCiphered {
		if char >= 'a' && char <= 'z' {
			char = char + 13
			if char > 'z' {
				char = char - 26
			}
		}
		fmt.Printf("%c", char)
	}
}
