package main

import "fmt"

func main() {
	var ConferenceName = "Go Conference"
	const ConferenceTickets = 50
	var RemainingTickets = 50

	fmt.Printf("Welcome User, to %v booking application\n", ConferenceName)
	fmt.Printf("total tickets are %v and remainig tickets are %v\n", ConferenceName, RemainingTickets)
	fmt.Println("Book your tickets here")

}
