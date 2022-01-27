package main

import (
	"fmt"
	"strings"
	"boot.it-app/shared"
)

const availableTickets uint = 10
var appName = "BookIt"
var appGoal = "Tickets"
var remainingTickets uint = 10
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		
		// user input
		userFirstName, userLastName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketsNumber := shared.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("Sorry, You have entered too short First Name or Last Name.\nTry it again:")
			continue
		} else if !isValidEmail {
			fmt.Println("Sorry, You have entered incorrect Email.\nTry it again:")
			continue
		} else if !isValidTicketsNumber {
			fmt.Printf("Sorry, just %v tickets available!\n Choose number between 0 and %v", remainingTickets)
			continue
		} else {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userFirstName + " " + userLastName + "<" + userEmail + "> booked tickets")
			
			fmt.Println("Thank you,", userFirstName, userLastName, "for booking", userTickets, "tickets!")
			fmt.Println("You will receive a confirmaiton email at", userEmail)
			fmt.Println(remainingTickets, "of", availableTickets, "tickets are still available")

			clientFirstNames := getFirstNames()

			fmt.Printf("The first names of bookings are: %v\n", clientFirstNames)

			if remainingTickets == 0 {
				fmt.Println("Sold out. C U later!")
				break
			}
		}
	}

}


// Functions
func greetUsers() {
	fmt.Printf("Welcome at %v!\n", appName)
	fmt.Println("Book", appGoal, "in this app!")
}

func getFirstNames() []string {
	clientFirstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		clientFirstNames = append(clientFirstNames, names[0])
	}

	return clientFirstNames
}

func getUserInput() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Please, enter your First Name:")
	fmt.Scan(&userFirstName)

	fmt.Println("Please, enter your Last Name:")
	fmt.Scan(&userLastName)

	fmt.Println("Please, enter your Email:")
	fmt.Scan(&userEmail)

	fmt.Println("Please, enter the number of tickets you want to book:")
	fmt.Scan(&userTickets)
	
	return userFirstName, userLastName, userEmail, userTickets
}