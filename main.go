package main

import "fmt"

func main() {
	conferenceName := "Go Conference 2023" 
	// alternative syntax for assigning variables
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	// uint is an unsigned integer, cannot be negative

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("Total tickets available ** %v ** and only ~ %v ~ tickets left!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Tickets Now!")	

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets with us!\n", firstName, lastName,userTickets)
	fmt.Printf("An email confirmation has been sent to %v.\n", email)

}