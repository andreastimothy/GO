package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-02/shared-projects/go-exercise-parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestAttPark(t *testing.T) {
	//AC1 Story 1
	t.Run("should return ticket when parking a car", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(5)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)
		ticket, err := attendant.AttPark(car)

		assert.NotNil(t, ticket)
		assert.Nil(t, err)
	})

	//AC5 Story 1
	t.Run("should not park if capacity is full", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(0)
		parkinglot2 := parking.NewParkingLot(0)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)
		ticket, err := attendant.AttPark(car)
		expectedError := parking.ParkingIsFullError{}

		assert.Nil(t, ticket)
		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
	})

	//AC6 Story 1
	t.Run("should not park when car already parked", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(5)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)
		ticket, err := attendant.AttPark(car)

		assert.NotNil(t, ticket)
		assert.Nil(t, err)

		car = parking.NewCar("B 1234 ABC")
		ticket, err = attendant.AttPark(car)
		expectedError := parking.CarAlreadyParkedError{}

		assert.Nil(t, ticket)
		assert.NotNil(t, err)
		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
	})
}

func TestAttUnpark(t *testing.T) {
	//AC2 Story 1
	t.Run("should return car when giving a valid ticket", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(5)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)
		ticket, err := attendant.AttPark(car)

		assert.Nil(t, err)

		returnedCar, err := attendant.AttUnpark(ticket)

		assert.Nil(t, err)
		assert.Same(t, car, returnedCar)
	})

	//AC3 Story 1
	t.Run("should return error when giving an invalid ticket", func(t *testing.T) {
		parkinglot1 := parking.NewParkingLot(5)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)
		ticket := parking.NewTicket("A")
		returnedCar, err := attendant.AttUnpark(ticket)
		expectedError := parking.InvalidTicketError{}

		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
		assert.Nil(t, returnedCar)
	})

	t.Run("should return error when unpark without a ticket", func(t *testing.T) {
		parkinglot1 := parking.NewParkingLot(5)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)
		returnedCar, err := attendant.AttUnpark(nil)
		expectedError := parking.InvalidTicketError{}

		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
		assert.Nil(t, returnedCar)
	})

	//AC4 Story 1
	t.Run("should not be able to unpark with the same ticket twice", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(5)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)
		ticket, _ := attendant.AttPark(car)
		returnedCar, err := attendant.AttUnpark(ticket)

		assert.Nil(t, err)
		assert.Same(t, car, returnedCar)

		returnedCar2, err2 := attendant.AttUnpark(ticket)
		expectedError := parking.InvalidTicketError{}
		
		assert.Exactly(t, expectedError, err2)
		assert.ErrorContains(t, err2, expectedError.Error())
		assert.Nil(t, returnedCar2)
	})
}

func TestAvailableLots(t *testing.T) {
	t.Run("should decreasing available lots if another lot is full", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(1)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)

		assert.Equal(t, 2, len(attendant.AvailableLots))

		attendant.AttPark(car)

		assert.Equal(t, 1, len(attendant.AvailableLots))
	})

	t.Run("should be able to park again after a car unparked", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(1)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)

		ticket, _ := attendant.AttPark(car)

		assert.Equal(t, 1, len(attendant.AvailableLots))

		attendant.AttUnpark(ticket)

		assert.Equal(t, 2, len(attendant.AvailableLots))

		attendant = parking.NewAttendant(parkinglot)
		car = parking.NewCar("B 1234 ABC")
		attendant.AttPark(car)

		assert.Equal(t, 1, len(attendant.AvailableLots))
	})
}

func TestStyle(t *testing.T) {
	t.Run("should park at lot with the most capacity", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(1)
		parkinglot2 := parking.NewParkingLot(5)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)

		attendant.NewStyle(&parking.ParkByMostCapacity{})

		attendant.AttPark(car)

		assert.Equal(t, 2, len(attendant.AvailableLots))
	})

	t.Run("should park sequentially by default", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(1)
		parkinglot2 := parking.NewParkingLot(5)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)

		attendant.AttPark(car)

		assert.Equal(t, 1, len(attendant.AvailableLots))
	})

	t.Run("should change style at anytime", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot1 := parking.NewParkingLot(1)
		parkinglot2 := parking.NewParkingLot(5)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)

		attendant.NewStyle(&parking.ParkByMostCapacity{})
		attendant.AttPark(car)

		assert.Equal(t, 2, len(attendant.AvailableLots))

		car = parking.NewCar("B 1234 ACD")
		parkinglot1 = parking.NewParkingLot(1)
		parkinglot2 = parking.NewParkingLot(5)
		parkinglot = []*parking.Lot{parkinglot1, parkinglot2}
		attendant = parking.NewAttendant(parkinglot)

		attendant.NewStyle(&parking.ParkSequentially{})
		attendant.AttPark(car)

		assert.Equal(t, 1, len(attendant.AvailableLots))
	})

	t.Run("should park at lot with most space", func(t *testing.T) {
		car1 := parking.NewCar("B 1234 ABC")
		car2 := parking.NewCar("B 1234 ABD")
		car3 := parking.NewCar("B 1234 ABE")
		car4 := parking.NewCar("B 1234 ABF")
		parkinglot1 := parking.NewParkingLot(4)
		parkinglot2 := parking.NewParkingLot(2)
		parkinglot := []*parking.Lot{parkinglot1, parkinglot2}
		attendant := parking.NewAttendant(parkinglot)

		attendant.AttPark(car1)
		attendant.AttPark(car2)
		attendant.AttPark(car3)

		attendant.NewStyle(&parking.ParkByMostSpace{})
		attendant.AttPark(car4)

		assert.Equal(t, 2, len(attendant.AvailableLots))

	})
}