package main

import (
	"fmt"
	"strings"
)

func main(){

	conferenceName := "Golands Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	//fmt.Printf("conference tickets is %T, remaining tickets is %T, conferenceName is %T \n", conferenceTickets,remainingTickets,conferenceName)

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Be sure to Go get your tickets!\n")

	var bookings []string

	for{
		
	

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask user for their name
		
		fmt.Printf("Enter your first name:\n")

		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: \n")

		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: \n")

		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: \n")

		fmt.Scan(&userTickets)

		remainingTickets= remainingTickets  - userTickets
		
		bookings = append(bookings, firstName + " " + lastName)
		
		fmt.Printf("Whole slice: %v \n", bookings)

		fmt.Printf("First value: %v \n", bookings[0])

		fmt.Printf("Slice is of type: %T \n", bookings )

		fmt.Printf("Slice length is %v \n", len(bookings))

		fmt.Printf("User %v %v booked %v tickets. \n", firstName,lastName,userTickets)

		fmt.Printf("%v tickets remain for the %v \n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking :=range bookings {
			var names = strings.Fields(booking)

			fmt.Print(names[0])
			
			firstNames = append(firstNames, names[0])

		}

		fmt.Printf("\nFirst names of the bookings are: %v\n", firstNames)
		
		noTicketsRemaining := remainingTickets ==0

		if noTicketsRemaining {
			//end program
			fmt.Println("Conference fully booked. ")
			break
		}
		
	}

	


	
}

