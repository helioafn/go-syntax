// Listing 9.9 - Decoding runes
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}

	for _, c := range question {
		fmt.Printf("%c ", c)
	}

	// Quick check 9.6
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("\n", utf8.RuneCountInString(alphabet), "runes") // 26 runes
	fmt.Println(len(alphabet), "bytes")                          // 26 bytes

	// 9.6.2 A: 2 bytes in rune '¿'

	/*
		Summary
		* Strings use a variable length encoding called UTF-8, where each character consumes 1–4 bytes.
		* The range keyword can decode a UTF-8 encoded string into runes.
	*/

}
