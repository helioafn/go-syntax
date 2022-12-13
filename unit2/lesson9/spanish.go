// Listing 9.8 - The utf8 package
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")                    // 15 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes") // 12 runes

	c, size := utf8.DecodeRuneInString(question)     // Returns the first character and the number of bytes consumed
	fmt.Printf("First rune: %c %v bytes\n", c, size) // First rune: ¿ 2 bytes

	myName := "Hélio Augusto Ferreira Nunes"
	fmt.Println(len(myName), "bytes")                    // 29 bytes
	fmt.Println(utf8.RuneCountInString(myName), "runes") // 28 runes 'é' is a single runec
	ch, s := utf8.DecodeRuneInString(myName)
	fmt.Printf("First rune: %c %v bytes", ch, s)

}
