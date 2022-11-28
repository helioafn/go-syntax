// Lesson 5

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const distanceToMars = 62_100_000 // Considering departure date Oct. 13, 2020
	const hoursInADay = 24
	const minFlightSpeed = 16 // km/s
	const maxFlightSpeed = 30 // km/s

	fmt.Println("Spaceline        Days Trip type Price")
	fmt.Println("=====================================")

	flightCompany := ""

	// Generating 10 tickets
	for i := 0; i < 10; i++ {

		// Calculating the flight speed
		flightSpeed := rand.Intn(maxFlightSpeed-minFlightSpeed+1) + minFlightSpeed // [0, 14[ + 16

		// flightSpeed is on km/s, that's why it has ...60 / 60 on the evaluation
		flightDuration := distanceToMars / flightSpeed / 60 / 60 / hoursInADay // Days
		// fmt.Println("Days:", flightDuration)

		flightPrice := 0
		// fmt.Println(flightPrice)

		// Choosing if it's round trip or not
		isRoundTrip := false // "Default" values
		flightType := "One-way"
		if rt := rand.Intn(2); rt == 0 {
			isRoundTrip = true
			flightType = "Round-trip"
		}

		if flightSpeed > 23 && isRoundTrip {
			flightPrice = (flightSpeed * 2) + 14
		} else if flightSpeed > 23 {
			flightPrice = flightSpeed + 14
		} else {
			flightPrice = flightSpeed + 14
		}

		// fmt.Println("Travel price:", flightPrice)
		// fmt.Println("The trip is round trip?", isRoundTrip)

		// Randomly choosing the flight company
		switch flightCompanyChoice := rand.Intn(3); flightCompanyChoice {
		case 0:
			flightCompany = "Virgin Galactic"
		case 1:
			flightCompany = "Space Adventures"
		case 2:
			flightCompany = "SpaceX"
		}

		// fmt.Println(flightCompany)
		// Priting ticket
		fmt.Printf("%-17v %v %-11v $%3v\n", flightCompany, flightDuration, flightType, flightPrice)
	}
}
