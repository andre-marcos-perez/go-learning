package main

import "fmt"

func main() {

    var conferenceName string = "Go Conference"
    const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Println("He have a total of", conferenceTickets, "ticket and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string

	var userTickets int

	fmt.Println("Enter you first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter you last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number if tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
