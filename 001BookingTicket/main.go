package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstName string
	var lastName string
	var Email string
	var tickets int
	var booking []string

	totalTickets := 50
	remainingTicket := 50

	fmt.Printf("Total Tickets are %v and Remaining Tickets are %v\n", totalTickets, remainingTicket)
	for{
		fmt.Println("Enter you FirstName")
		fmt.Scan(&firstName)

		fmt.Println("Enter you LastName")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter you email address")
		fmt.Scan(&Email)
		
		fmt.Println("Enter number of tickets you want")
		fmt.Scan(&tickets)

		isValidName := len(firstName)>3 && len(lastName)>3
		isValidEmail := strings.Contains(Email,"@")
		isValidTicket := tickets <= remainingTicket && tickets>0

		if isValidEmail && isValidName && isValidTicket{
			remainingTicket = remainingTicket- tickets
	
			booking = append(booking, firstName+" "+lastName)
			
			arrFirstName := []string{}
			for _, booking := range booking{
				var name = strings.Fields(booking)
				arrFirstName = append(arrFirstName, name[0])
			}
	
			fmt.Println(arrFirstName)
	
			fmt.Printf("Thanks %v %v for booking %v tickets. You'll get a confirmation at %v\n", firstName, lastName, tickets, Email)
			
			fmt.Printf("Total Tickets are %v and Remaining Tickets are %v\n", totalTickets, remainingTicket)
			
		}else if remainingTicket == 0{
			fmt.Println("House Full, Booking Closed")
			break
		}else{
			if !isValidName{
				fmt.Printf("Your name %v is too short\n",firstName)
			}
			if !isValidEmail{
				fmt.Printf("Your email %v doesnot have @\n",Email)
			}
			if !isValidTicket{
				fmt.Printf("We have %v remaining tickets. You have entered %v tickets.\nRequest Denied\n", remainingTicket, tickets)
			}


		}
	}
}