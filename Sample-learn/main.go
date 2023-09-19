// Here we create a booking app for a concert

package main

import "fmt"

func main() {
	var concertName = "Mamma mia"
	const concertTickets = 50
	var remainingTickets = 50

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

	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)
}
