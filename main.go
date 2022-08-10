package main

import (
	"fmt"
	"strings"
)


func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}


	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)


	fmt.Printf("Welcome to %v booking application application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	
	// looping it all up
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask for user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		// appending bookings to the bookings array
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("the whole array: %v\n", bookings)
		fmt.Printf("the first value: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))

		fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)


		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These first names of bookings are : %v\n", firstNames)
		
		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out come back next year.")
			break
		}
		
	}
	
}