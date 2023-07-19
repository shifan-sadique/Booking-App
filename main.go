package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	var firstname string
	var lastname string
	var userTickets int
	var email string
	var bookings []string

	for {
		greetUser(conferenceName, conferenceTickets, remainingTickets)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstname)

		fmt.Println("enter your last name")
		fmt.Scan(&lastname)

		fmt.Println("enter your email id")
		fmt.Scan(&email)

		fmt.Println("enter the number of tickets required")
		fmt.Scan(&userTickets)

		isValidName := len(firstname) >= 2 && len(lastname) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstname+" "+lastname)
			fmt.Printf("type is %T", bookings)

			fmt.Printf("Thank You %v %v for booking %v tickets. You'll have your confirmation at %v. \n", firstname, lastname, userTickets, email)
			fmt.Println("The remaining tickets available for booking are", remainingTickets)
			var firstname = printFirstNames(bookings)
			fmt.Printf("these are all our bookings: %v \n", firstname)

			if remainingTickets == 0 {
				fmt.Println("The tickets are sold out")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email entered doesnt contain a @ symbol")
			}
			if !isValidTicketNumber {

				fmt.Println("Number of tickets invalid")
			}
		}

	}
}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Welcome to our %v \n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets")
	fmt.Println("There are", remainingTickets, "remaining for booking")
	fmt.Println("Book your tickets here")
}

func printFirstNames(bookings []string) []string {
	firstnames := []string{}
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	return firstnames
}
