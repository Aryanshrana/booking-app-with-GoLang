package main

import "fmt"

func main() {
	var ConferenceName = "Go Conference"
	const ConferenceTickets = 50
	var RemainingTickets = 50

	//alternative way of declaring a variable and assigning a variable to it
	//with this you can not declare constants

	Conferencelocation := "Jaipur"
	fmt.Printf("the conference is at %v", Conferencelocation)

	//printing the datatype of variable
	fmt.Printf("ConferenceTickets is %T and ConferenceName is %T\n", ConferenceTickets, ConferenceName)

	fmt.Printf("Welcome User, to %v booking application\n", ConferenceName)
	fmt.Printf("total tickets are %v and remainig tickets are %v\n", ConferenceTickets, RemainingTickets)
	fmt.Println("Book your tickets here")

	//asking user name when not assiging a value to variable at the time of declaration you have to explicitly tell go the datatype like below
	var userName string
	userName = "Aryansh"

	var userTicket int
	userTicket = 2

	fmt.Printf("user %v booked %v tickets", userName, userTicket)

}
