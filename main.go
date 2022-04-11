package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferencename = "sirConference"
	const conferencetickets = 50
	//int that cant be negative
	var remainingTickets uint = 50
	//ARRAY just add a size in the [] to make it an array
	var bookings []string //this is a slice its also the same as 'bookings := []string{}

	//for condition
	for //Remainingtickets > 0 && len(bookings) < 50
	{
		greetUsers(conferencename, conferencetickets, int(remainingTickets))
		//function get user input

		var username string
		var lastname string
		var userTickets uint
		var email string

		//user input

		fmt.Print("enter first  name: ")
		fmt.Scanln(&username)

		fmt.Print("enter last name: ")
		fmt.Scanln(&lastname)

		fmt.Print("enter Email. Address: ")
		fmt.Scanln(&email)

		fmt.Print("how many tickets? ")
		fmt.Scanln(&userTickets)

		//fmt.Println("Enter email address")
		//fmt.Scanln(&email)
		var ticketsBooked = userTickets

		//call validate user input function
		isValidName, isValidEmail, isValidTicketNumber := validateUserinput(username, lastname, email, userTickets, remainingTickets)
		//stop from proceeding to next step if ticket number is invalid
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(ticketsBooked)
			//fmt.Println("Remaining tickets:", Remainingtickets)

			//for loop iteration
			//two values for each iteration which is the _ and the element
			//_ is blank identifier
			//element is booking; bookings is the one being iterated; range is the one that allows us to iterate over elements for different data structures
			//range for arrays and slices it gives back the index and value for each element

			//adding a new element in the next slice
			bookings = append(bookings, username+" "+lastname) //add element to the next index

			//get function print first names
			firstnames := getFirstNames(bookings)
			fmt.Printf("users first name%v \n", firstnames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name, try again")
			}
			if !isValidEmail {
				fmt.Println("Invalid email, try again")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number, try again")
			}
			// continue - causes the loop to skip remainder of its body and repeating the 1st step
			//continue
		}
	}
	//bookings = append(bookings, "tite")
	//show all array
	//fmt.Println(bookings)
	//show array[0]
	//fmt.Println(bookings)
	//array type (string)
	//fmt.Printf("Array type: %T\n", bookings)
	//array length
	//fmt.Println(len(bookings))

	/* Switch Statement

	var city string = "London"
	switch city {
	case "New York":
		//exec code

	//if same code
	case "Singapore", "Myanmar":
		//exec code
	case "London":
		//exec code
	default:
		fmt.Println("No valid city selected")
	}*/
}

//encapsulated functions
func greetUsers(confName string, confTickets int, confRemaining int) {
	fmt.Printf("Welcome %v users \n", confName)
	fmt.Printf("total no. tickets: %v \nremaining tickets %v \n", confTickets, confRemaining)
}

//get name function
func getFirstNames(bookings []string) []string {
	firstnames := []string{}
	for _, booking := range bookings {
		//The Fields() function takes the string as a parameter, which we need to splits.
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	return firstnames
}

func validateUserinput(username string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	//check if name length is more than 2 characters
	var isValidName bool = len(username) >= 2 && len(lastname) >= 2
	//check if email is valid
	var isValidEmail bool = strings.Contains(email, "@")
	//check if ticket is valid
	var isValidTicketNumber bool = userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
