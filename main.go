package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferencename = "sirConference"
	const conferencetickets = 50
	//int that cant be negative
	var Remainingtickets uint = 50
	//ARRAY just add a size in the [] to make it an array
	var bookings []string //this is a slice its also the same as 'bookings := []string{}
	for {

		var username string
		var lastname string
		var userTickets uint

		fmt.Printf("welcome to %v booking application\n", conferencename)
		fmt.Printf("total no. tickets: %v \nremaining tickets %v \n", conferencetickets, Remainingtickets)
		fmt.Print("enter name: ")
		//user input
		fmt.Scanln(&username)
		fmt.Print("enter last name: ")
		fmt.Scanln(&lastname)
		fmt.Print("how many tickets? ")
		fmt.Scanln(&userTickets)
		var ticketsBooked = userTickets

		//stop from proceeding to next step if ticket number is invalid
		if userTickets > Remainingtickets {
			fmt.Printf("we only have %v tickets left you can't book %v tickets exceeding %v tickets", Remainingtickets, userTickets, Remainingtickets)
			// causes the loop to skip remainder of its body and repeating the 1st step
			continue
		}
		Remainingtickets = Remainingtickets - uint(ticketsBooked)
		fmt.Println("Remaining tickets:", Remainingtickets)

		firstnames := []string{}
		//for loop iteration
		//two values for each iteration which is the _ and the element
		//_ is blank identifier
		//element is booking; bookings is the one being iterated; range is the one that allows us to iterate over elements for different data structures
		//range for arrays and slices it gives back the index and value for each element

		bookings = append(bookings, username+" "+lastname) //add element to the next index

		for _, booking := range bookings {
			//The Fields() function takes the string as a parameter, which we need to splits.
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}
		fmt.Printf("users first name%v \n", firstnames)

		if Remainingtickets == 0 {
			fmt.Println("Tickets are sold out")
			break
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

	//adding a new element in the next slice
}
