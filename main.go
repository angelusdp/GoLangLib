package main

import (
	"fmt"
	"sync"
	"time"
)

//package variables
const conferencetickets = 50

var conferencename = "sirConference"

//int that cant be negative
var remainingTickets uint = 50

//ARRAY just add a size in the [] to make it an array
var bookings = make([]UserData, 0)

//var bookings = make([]map[string]string, 0) //this is a slice its also the same as 'bookings := []string{} //made it a slice of maps

type UserData struct {
	username        string
	lastname        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//for condition
	//for //Remainingtickets > 0 && len(bookings) < 50
	//{
	greetUsers()
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

	//call validate user input function
	isValidName, isValidEmail, isValidTicketNumber := validateUserinput(username, lastname, email, userTickets)
	//stop from proceeding to next step if ticket number is invalid
	if isValidName && isValidEmail && isValidTicketNumber {
		//bookTickets(remainingTickets, userTickets, bookings, username, lastname)
		bookTickets(userTickets, username, lastname, email)
		wg.Add(1)
		//call send ticket
		go sendTickets(userTickets, username, lastname, email) //make concurrent thread using go keyword
		//get function print first names
		firstnames := getFirstNames()
		fmt.Printf("users first name%v \n", firstnames)

		if remainingTickets == 0 {
			fmt.Println("Tickets are sold out")
			//break
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
	wg.Wait() //wait for thread to be done b4 the app can exit
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
//}

//encapsulated functions
func greetUsers() {
	fmt.Printf("Welcome %v users \n", conferencename)
	fmt.Printf("total no. tickets: %v \nremaining tickets %v \n", conferencetickets, remainingTickets)
}

//get name function
func getFirstNames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		//The Fields() function takes the string as a parameter, which we need to splits.
		//var names = strings.Fields(booking) //local variable
		firstnames = append(firstnames, booking.username) //get struct username
	}
	return firstnames
}

func bookTickets(userTickets uint, username string, lastname string, email string) {
	remainingTickets = remainingTickets - userTickets
	fmt.Println("Remaining tickets:", remainingTickets)
	//for loop iteration
	//two values for each iteration which is the _ and the element
	//_ is blank identifier
	//element is booking; bookings is the one being iterated; range is the one that allows us to iterate over elements for different data structures
	//range for arrays and slices it gives back the index and value for each element

	//create a userdata struct object
	var userData = UserData{
		username:        username,
		lastname:        lastname,
		email:           email,
		numberOfTickets: userTickets,
	}
	//var userData = make(map[string]string) // create a map
	//userData["first Name"] = username //add  "first name" and data to map [0]

	//adding a new element in the next slice
	bookings = append(bookings, userData) //add element to the next index
	fmt.Printf("List of bookings is %v \n", bookings)
	fmt.Printf("Thank you for booking: %v, an email will be sent to your email %v \n", bookings, email)
}

func sendTickets(userTickets uint, username string, lastname string, email string) {
	//give the program to process using sleep function 10 sec delay
	time.Sleep(10 * time.Second) //block execution of thread for defined times
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, username, lastname)
	fmt.Println("###########")
	fmt.Printf("sending: \n %v tickets \n to: \n%v \n", ticket, email)
	fmt.Println("###########")
	wg.Done() //removes the thread wg.add(1)
}
