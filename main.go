package main

import (
	"fmt"
	"go-booking-app/helper"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Con"
var bookings = make([]UserData, 0)
var remainingTickets uint = conferenceTickets

// Ensure all threads close before closing app
var wg = sync.WaitGroup{}

type UserData struct {
	firstName    string
	lastName     string
	email        string
	ticketsCount uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("Sorry, your first or last name was too short. Please use at least 2 characters.")
			// Skips to next iteration of loop
			continue
		}

		if !isValidEmail {
			fmt.Println("Your email was missing '@'")
			continue
		}

		if !isValidUserTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining. Please try again.\n", remainingTickets)
			continue
		}

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		var user = createBooking(firstName, lastName, email, userTickets)
		bookings = append(bookings, user)
		wg.Add(1)
		go sendTicket(user)

		fmt.Printf("Ticket purchasers: %v\n", getFirstNames())
		fmt.Println(bookings)

		if remainingTickets == 0 {
			fmt.Println("Conference is sold out! Come back next year.")
			break
		}
	}

	wg.Wait()
}

func createBooking(firstName string, lastName string, email string, userTickets uint) UserData {
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["ticketsCount"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		ticketsCount: userTickets,
	}

	return userData
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking app.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets. %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

// return firstName string, lastName string, email string, userTickets uint
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name.")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name.")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email.")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(user UserData) {
	// Wait 10 seconds to send
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", user.ticketsCount, user.firstName, user.lastName)

	fmt.Println("##########################")
	fmt.Printf("Sending ticket to:\n %v\nto email address %v\n", ticket, user.email)
	fmt.Println("##########################")

	wg.Done()
}

// Notes
// This won't work
// const conferenceTickets := 50
// var bookings = [50]string{"Nana", "Nicole"}
// var bookings [50]string
// Use slice instead of array
// fmt.Printf("The whole array: %v\n", bookings)
// fmt.Printf("The first array value: %v\n", bookings[0])
// fmt.Printf("The array type: %T\n", bookings)
// fmt.Printf("The array length: %v\n", len(bookings))
// bookings[0] = firstName + lastName
// else if conditional {}
// for remainingTickets > 0 {

// must run every file go run main.go helper.go
// better go run .
// Captilize name of exported functions to share across packages
// Captilize variable names to share across packages

// Can't mix data types in a map
// var bookings = make([]map[string]string, 0)
// bookings[0]["firstName"]

// Concurrency is handled by adding 'go' in front of calls
// go runExpensiveFunction()
// if main thread stops, everything does
