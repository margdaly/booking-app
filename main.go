package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference 2023" 
	// alternative syntax for assigning variables
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	// uint is an unsigned integer, cannot be negative
	var bookings []string
	// Slice of strings, dynamic in size

	// var bookings = [50]string{}
	// array with fixed size of 50 elements and type of string

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("Total tickets available ** %v ** and only ~ %v ~ tickets left!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Tickets Now!")	

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	// ask user for their name, email and ticket amount
	// Scan is a blocking function, it will wait for user input
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	// append is a built-in function to add elements to a slice

	fmt.Printf("Thank you %v %v for booking %v tickets with us!\n", firstName, lastName,userTickets)
	fmt.Printf("An email confirmation has been sent to %v.\n", email)

	fmt.Printf("\nTotal tickets available ** %v ** and only ~ %v ~ tickets left!\n", conferenceTickets, remainingTickets)

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("People with bookings so far: %v\n", firstNames)

}