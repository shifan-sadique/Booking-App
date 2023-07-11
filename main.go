package main

import (
	"fmt"
)

func main()  {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	var firstname string
	var lastname string
	var userTickets int
	var email string
	var bookings [] string

	fmt.Println("Welcome to", conferenceName)
	fmt.Println("We have a total of",conferenceTickets,"tickets")
	fmt.Println("There are",remainingTickets,"remaining for booking")
	fmt.Println("Book your tickets here")

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("enter your last name")
	fmt.Scan(&lastname)

	fmt.Println("enter your email id")
	fmt.Scan(&email)

	fmt.Println("enter the number of tickets required")
	fmt.Scan(&userTickets)
	
	remainingTickets= remainingTickets-userTickets
	bookings =append(bookings,firstname +" "+ lastname ) 
	fmt.Printf("type is %T",bookings)

	fmt.Printf("Thank You %v %v for booking %v tickets. You'll have your confirmation at %v. \n",firstname,lastname,userTickets,email)
	fmt.Println("The remaining tickets available for booking are",remainingTickets)
	fmt.Println(bookings[0])

}
 