package main

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-02/shared-projects/go-exercise-parking-lot/parking"
	"bufio"
	"os"
)

func main() {

	parkinglot1 := parking.NewParkingLot(2)
	parkinglot2 := parking.NewParkingLot(5)
	parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
	attendant := parking.NewAttendant(parkinglot)
	ticketList := map[string]*parking.Ticket{}
	reader := bufio.NewReader(os.Stdin)
	a := true

	for a {
		var input int
		fmt.Println("==================================")
		fmt.Println("Welcome to our Car Parking System!")
		fmt.Println("Please choose 1 of this options:")
		fmt.Println("1. Park")
		fmt.Println("2. Unpark")
		fmt.Println("3. Exit")
		fmt.Println("==================================")
		fmt.Print("Input here: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Print("Write your plate number: ")
			plate, _ := reader.ReadString('\n')
			car := parking.NewCar(plate)
			ticket, err := attendant.AttPark(car)

			if err != nil {
				fmt.Println("")
				fmt.Println("Error:", err)
				fmt.Println("==================================")
				fmt.Println("")
				fmt.Println("")
			} else {
				ticketId := parking.GetTicket(ticket)
				ticketList[plate] = ticket
				fmt.Println("")
				fmt.Println("Parking done!")
				fmt.Print("Here's your ticket: ", ticketId)
				fmt.Println("==================================")
				fmt.Println("")
				fmt.Println("")
			}
		case 2:
			fmt.Print("Write your ticket: ")
			ticket, _ := reader.ReadString('\n')
			returnedCar, err := attendant.AttUnpark(ticketList[ticket])

			if err != nil {
				fmt.Println("")
				fmt.Println("Error:", err)
				fmt.Println("==================================")
				fmt.Println("")
				fmt.Println("")
			} else {
				plateNumber := parking.GetPlate(returnedCar)
				delete(ticketList, ticket)
				fmt.Println("")
				fmt.Println("Unparking done!")
				fmt.Print("Car's plate number: ", plateNumber)
				fmt.Println("==================================")
				fmt.Println("")
				fmt.Println("")
			}
		case 3:
			fmt.Println("Thank you for using our service!")
			fmt.Println("==================================")
			fmt.Println("")
			a = false
		}
	}
}