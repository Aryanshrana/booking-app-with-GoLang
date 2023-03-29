package main

import (
	"fmt"
	"strings"
)

const ConferenceTickets = 50

var ConferenceName = "Go Conference"
var RemainingTickets uint = 50
var bookings = []string{}

func main() {

	//alternative way of declaring a variable and assigning a variable to it
	//with this you can not declare constants

	//Conferencelocation := "Jaipur"
	//fmt.Printf("the conference is at %v\n", Conferencelocation)

	//printing the datatype of variable
	//fmt.Printf("ConferenceTickets is %T and ConferenceName is %T\n", ConferenceTickets, ConferenceName)
	greetUser()

	//asking user name when not assiging a value to variable at the time of declaration you have to explicitly tell go the datatype like below

	//array in go -assign value at declaration
	//var bookings = [50]string{"Aryansh","Shubh","Henry","Peter"}
	//declaring first assigning after
	//var bookings [50]string - > array
	// slices --> it is more dynamic

	for {

		if RemainingTickets == 0 {
			fmt.Println("All tickets are booked, you are late")
			break
		}
		//getting user input
		firstName, lastName, email, userTicket := getUserInput()

		//checking user validation
		isvalidName, isvalidEmail, isvalidTickets := validateUserInput(firstName, lastName, email, userTicket)

		if isvalidName && isvalidEmail && isvalidTickets {
			//bookTickets
			bookTickets(userTicket, firstName, lastName, email)

			firstnames := getfirstnames()
			fmt.Printf("The firstnames of all the persons who booked a ticket %v\n", firstnames)

		} else {
			if !isvalidName {
				fmt.Println("User input name is invalid,They are very short")
			}
			if !isvalidEmail {
				fmt.Println("User input email is invalid, It does not have @")
			}
			if !isvalidTickets {
				fmt.Println("User Input tickets are invalid")
			}
			fmt.Println("Try again!")
			//fmt.Println("Your input Data is invalid, Try again!")
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome User, to %v booking application\n", ConferenceName)
	fmt.Printf("total tickets are %v and remainig tickets are %v\n", ConferenceTickets, RemainingTickets)
	fmt.Println("Book your tickets here to attend")
}

func getfirstnames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		// it expect two values from bookings , so we have to use two variables to coolece thos values but we are not using first value that is index , so instead of using index we can use underscore
		// _ are used to represent those values which we don't want to use
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	return firstnames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	//asking user input
	fmt.Println("enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTickets(userTicket uint, firstName string, lastName string, email string) {
	RemainingTickets = RemainingTickets - userTicket
	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, ","+firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve notifications regarding conference on you email %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("now %v tickets remaining for %v\n", RemainingTickets, ConferenceName)
}
