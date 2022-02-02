package main

import (
	"fmt"
//	"strconv"
	"boot.it-app/shared"
	"time"
)

// struct (class)
type OtherUserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

const availableTickets uint = 10
var appName = "BookIt"
var appGoal = "Tickets"
var remainingTickets uint = 10
//var bookings = make([]map[string]string, 0)
var bookings = make([]OtherUserData, 0)



func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		
		// user input
		userFirstName, userLastName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketsNumber := shared.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("Sorry, You have entered too short First Name or Last Name.\nTry it again:\n")
			continue
		} else if !isValidEmail {
			fmt.Println("Sorry, You have entered incorrect Email.\nTry it again:\n")
			continue
		} else if !isValidTicketsNumber {
			fmt.Printf("Sorry, just %v tickets available!\n Choose number between 0 and %v\n", remainingTickets)
			continue
		} else {
			remainingTickets = remainingTickets - userTickets

			// a map
			/*
			var userData = make(map[string]string)
			userData["userFirstName"] = userFirstName
			userData["userLastName"] = userLastName
			userData["userEmail"] = userEmail
			userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
			*/

			var otherUserData = OtherUserData {
				firstName: userFirstName,
				lastName: userLastName,
				email: userEmail,
				numberOfTickets: userTickets,
			}
			

			bookings = append(bookings, otherUserData)
			fmt.Printf("List of bookings is: %v\n", bookings)

			// WaitGroup can be added for waiting when sending happens before app ends.
			go sendTicket(userTickets, userFirstName, userLastName, userEmail)
			
			fmt.Println("Thank you,", userFirstName, userLastName, "for booking", userTickets, "tickets!")
			fmt.Println("You will receive a confirmaiton email at", userEmail)
			fmt.Println(remainingTickets, "of", availableTickets, "tickets are still available")

			//clientFirstNames := getFirstNames()

			//fmt.Printf("The first names of bookings are: %v\n", clientFirstNames)

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
		//clientFirstNames = append(clientFirstNames, booking["firstName"])
		clientFirstNames = append(clientFirstNames, booking.firstName)
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

func sendTicket(userTickets uint, userFirstName string, userLastName string, userEmail string) {
	time.Sleep(3 * time.Second)

	//fmt.Printf("%v tickets for %v %v booked", userTickets, userFirstName, userLastName)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, userFirstName, userLastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address: %v\n", ticket, userEmail)
	fmt.Println("###############")
}