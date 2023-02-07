package main

import (
	"fmt"
	"strings"
)

func main() {

	var goConference = "Raghu Go Conference"
	const totalTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v Application \n", goConference)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", totalTickets, remainingTickets)
	fmt.Printf("Get your Tickets to attend \n")
	for {
		var firstName string
		var lastName string
		var userTickets int
		var email string
		bookings := []string{}
		fmt.Printf("Enter Your First Name:")
		fmt.Scan(&firstName)
		fmt.Printf("Enter Your Last Name:")
		fmt.Scan(&lastName)
		fmt.Printf("Enter No Of Tickets:")
		fmt.Scan(&userTickets)

		fmt.Printf("Enter your email:")
		fmt.Scan(&email)

		if userTickets > remainingTickets {
			fmt.Println("we only %v tickets remaining, so you  can't book %v\n", remainingTickets, userTickets)
			break
		}
		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		// fmt.Printf("the whole array%v\n", bookings)
		// fmt.Printf("the first array%v\n", bookings[0])
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thank You %v %v booked %v for booking tickets. you will Receive confirmation at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, goConference)

		firstNam := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)

			firstNam = append(firstNam, names[0])
		}
		fmt.Printf("the firstNam %v\n", firstNam)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. come back next Year")
			break
		}
	}
}
