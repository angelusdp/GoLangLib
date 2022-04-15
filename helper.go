package main

import "strings"

/*global variable
var Hatdog = "some value" //capital 1st letter
*/
func validateUserinput(username string, lastname string, email string, userTickets uint) (bool, bool, bool) {
	//check if name length is more than 2 characters
	var isValidName bool = len(username) >= 2 && len(lastname) >= 2
	//check if email is valid
	var isValidEmail bool = strings.Contains(email, "@")
	//check if ticket is valid
	var isValidTicketNumber bool = userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
