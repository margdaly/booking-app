package main

import "fmt"

func main() {
	var conferenceName = "Go Conference 2023"
	const conferenceTickets = 50
	var remainingTickets = 50 

	fmt.Println("Welcome to", conferenceName, "booking app!")
	fmt.Println("Total tickets available *",conferenceTickets,"*")
	fmt.Println("And only *",remainingTickets,"* tickets left!")
	fmt.Println("Get Your Tickets Now!")
}
