package main

import "fmt"

func main() {
	var ConferenceName = "Go Conference"
	const ConferenceTickets = 50
	var RemainingTickets uint = 50

	//alternative way of declaring a variable and assigning a variable to it
	//with this you can not declare constants

	//Conferencelocation := "Jaipur"
	//fmt.Printf("the conference is at %v\n", Conferencelocation)

	//printing the datatype of variable
	//fmt.Printf("ConferenceTickets is %T and ConferenceName is %T\n", ConferenceTickets, ConferenceName)

	fmt.Printf("Welcome User, to %v booking application\n", ConferenceName)
	fmt.Printf("total tickets are %v and remainig tickets are %v\n", ConferenceTickets, RemainingTickets)
	fmt.Println("Book your tickets here")

	//asking user name when not assiging a value to variable at the time of declaration you have to explicitly tell go the datatype like below
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	//array in go -assign value at declaration
	//var bookings = [50]string{"Aryansh","Shubh","Henry","Peter"}
	//declaring first assigning after
	var bookings [50]string
	//asking user input
	fmt.Println("enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets: ")
	fmt.Scan(&userTicket)

	RemainingTickets = RemainingTickets - userTicket
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Persons who booked: %v\n", bookings)
	//printing length of array
	fmt.Printf("length of array- %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve notifications regarding conference on you email %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("now %v tickets remaining for %v", RemainingTickets, ConferenceName)

}
