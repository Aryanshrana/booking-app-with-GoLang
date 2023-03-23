package main

import (
	"fmt"
	"strings"
)

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
	//var bookings [50]string - > array
	// slices --> it is more dynamic
	var bookings []string

	for {

		if RemainingTickets == 0 {
			fmt.Println("All tickets are booked, you are late")
			break
		}

		//asking user input
		fmt.Println("enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets: ")
		fmt.Scan(&userTicket)

		if userTicket > RemainingTickets {
			fmt.Printf("Sorry we have only %v tickets available, you can not book %v tickets so book only in range\n", RemainingTickets, userTicket)
			continue
		}

		RemainingTickets = RemainingTickets - userTicket
		//bookings[0] = firstName + " " + lastName
		bookings = append(bookings, ","+firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve notifications regarding conference on you email %v\n", firstName, lastName, userTicket, email)
		fmt.Printf("now %v tickets remaining for %v\n", RemainingTickets, ConferenceName)

		//printing length of array
		//fmt.Printf("Length of slice- %v\n", len(bookings))
		firstnames := []string{}
		for _, booking := range bookings {
			// it expect two values from bookings , so we have to use two variables to coolece thos values but we are not using first value that is index , so instead of using index we can use underscore
			// _ are used to represent those values which we don't want to use
			var names = strings.Fields(booking)

			firstnames = append(firstnames, names[0])
		}

		fmt.Printf("The firstnames of all the persons who booked a ticket %v\n", firstnames)

	}

}
