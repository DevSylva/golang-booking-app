package helper

import (
	"fmt"
	"strings"
	"helper"
)

var conferenceName = "Go conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = []string{}



func main() {
	
	greetUsers()

	// looping it all up
	for {

		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateuserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail{
			// book ticket
			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			

			// call function print first names
			
			fmt.Printf("the first names of bookins are %v", getFirstNames(bookings))
			
			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out come back next year.")
				break
			}
		}else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of ticket you entered is invalid")
			}
		}
	}
	
}

func greetUsers() {
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("------------------------------")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	// fmt.Printf("These first names of bookings are : %v\n", firstNames)
	return firstNames
}



func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, email string, lastName string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}