// Here we create a booking app for a concert

package main

import "fmt"

var remainingTickets int
var concertTickets int
var concertName string

func main() {
	concertName = "Mamma mia"
	concertTickets = 50
	remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", concertName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", concertTickets, remainingTickets)
	fmt.Println("Get your tickets to attend.")

	// get username
	var userName string
	var userTicket int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter how many tickets you need: ")
	fmt.Scan(&userTicket)

	if userTicket > remainingTickets {
		fmt.Printf("Unfortunately we only have %v tickets left!!", remainingTickets)

	} else {
		fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)
		buyTicket(userTicket)
		fmt.Printf("We have %v tickets left", remainingTickets)
	}

}

func buyTicket(userTicket int) {
	remainingTickets = concertTickets - userTicket
}
