// Listing 10.5 - Converting a Boolean to a string
package main

import "fmt"

func main() {
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText) // Ready for launch: false

	yesNo := ""
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}

	fmt.Println("Ready for launch:", yesNo) // Ready for launch: no
}
