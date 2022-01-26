package main

import (
	"fmt"
	"strings"
)

func main() {
	appName := "BookIt"
	appGoal := "Tickets"

	const availableTickets uint = 50
	var remainingTickets uint = 50
	
	fmt.Printf("Welcome to the %v application!\n", appName)
	fmt.Println("Book", appGoal, "in this app!")

	bookings := []string{}

	for {
		var userFirstName string
		var userLastName string
		var userEmail string
		var numberOfTickets uint

		fmt.Println("Please, enter your First Name, Last Name and Email:")
		fmt.Println(">>> Use Enter button on your keyboard after each value! <<<")
		fmt.Scan(&userFirstName, &userLastName, &userEmail)
		fmt.Println("Please, enter the number of tickets you want to book:")
		fmt.Scan(&numberOfTickets)
		if numberOfTickets > remainingTickets {
			fmt.Printf("Sorry, just %v tickets available!\n", remainingTickets)
			break
		} else {
			remainingTickets = remainingTickets - numberOfTickets
			bookings = append(bookings, userFirstName + " " + userLastName + "<" + userEmail + "> booked tickets")
			
			fmt.Println("Thank you,", userFirstName, userLastName, "for booking", numberOfTickets, "tickets!")
			fmt.Println("You will receive a confirmaiton email at", userEmail)
			fmt.Println(remainingTickets, "of", availableTickets, "tickets are still available")

			clientFirstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				clientFirstNames = append(clientFirstNames, names[0])
			}

			fmt.Printf("These are all our bookings:\n %v\n", clientFirstNames)
		}
	}

	
}
