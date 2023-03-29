package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTickets := userTicket > 0 && userTicket <= RemainingTickets

	// --most important-- //
	//in go you can return any number of values simultaneously
	return isvalidName, isvalidEmail, isvalidTickets
}
