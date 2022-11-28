// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	// Listing 2.3
	const lightspeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println(distance/lightspeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightspeed, "seconds")

	// Quick check 2.3.1
	const travelSpeed = 100800 // km/h
	distance = 96300000        // km
	const earthHoursPerDay = 24

	fmt.Println(distance/travelSpeed/earthHoursPerDay, "days")
	// Quick check 2.3.2 -> const keyword

	// 2.4 Taking a shortcut
	// You can declare each variable on its own line
	// var distanceToYourHeart = 1337

	// Or declare them as a group
	// var (
	// 	runningCourseDistance = 2.3 // km
	// 	speed                 = 10  // km/h
	// )

	// Another option is to declare multiple variables on a single line
	// var distanceToMoon, speed = 1234567, 100150

	// Quick check 2.4
	// const hoursPerDay, minutesPerHour = 24, 60
}
