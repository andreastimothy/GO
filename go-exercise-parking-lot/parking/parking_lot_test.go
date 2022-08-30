package parking_test

import (
	"testing"
	"git.garena.com/sea-labs-id/batch-02/shared-projects/go-exercise-parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestPark(t *testing.T) {
	//AC1 Story 1
	t.Run("should return ticket when parking a car", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot := parking.NewParkingLot(10)
		ticket, err := parkinglot.Park(car)

		assert.NotNil(t, ticket)
		assert.Nil(t, err)
	})

	//AC5 Story 1
	t.Run("should not park if capacity is full", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot := parking.NewParkingLot(0)
		ticket, err := parkinglot.Park(car)
		expectedError := parking.ParkingIsFullError{}

		assert.Nil(t, ticket)
		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
	})

	//AC6 Story 1
	t.Run("should not park when car already parked", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot := parking.NewParkingLot(10)
		ticket, err := parkinglot.Park(car)

		assert.NotNil(t, ticket)
		assert.Nil(t, err)

		ticket, err = parkinglot.Park(car)
		expectedError := parking.CarAlreadyParkedError{}

		assert.Nil(t, ticket)
		assert.NotNil(t, err)
		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
	})
}

func TestUnpark(t *testing.T) {
	//AC2 Story 1
	t.Run("should return car when giving a valid ticket", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot := parking.NewParkingLot(10)
		ticket, err := parkinglot.Park(car)

		assert.Nil(t, err)

		returnedCar, err := parkinglot.Unpark(ticket)

		assert.Nil(t, err)
		assert.Same(t, car, returnedCar)
	})

	//AC3 Story 1
	t.Run("should return error when giving an invalid ticket", func(t *testing.T) {
		parkinglot := parking.NewParkingLot(10)
		ticket := parking.NewTicket("A")
		returnedCar, err := parkinglot.Unpark(ticket)
		expectedError := parking.InvalidTicketError{}

		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
		assert.Nil(t, returnedCar)
	})

	t.Run("should return error when unpark without a ticket", func(t *testing.T) {
		parkinglot := parking.NewParkingLot(10)
		returnedCar, err := parkinglot.Unpark(nil)
		expectedError := parking.InvalidTicketError{}

		assert.Exactly(t, expectedError, err)
		assert.ErrorContains(t, err, expectedError.Error())
		assert.Nil(t, returnedCar)
	})

	//AC4 Story 1
	t.Run("should not be able to unpark with the same ticket twice", func(t *testing.T) {
		car := parking.NewCar("B 1234 ABC")
		parkinglot := parking.NewParkingLot(10)
		ticket, _ := parkinglot.Park(car)
		returnedCar, err := parkinglot.Unpark(ticket)

		assert.Nil(t, err)
		assert.Same(t, car, returnedCar)

		returnedCar2, err2 := parkinglot.Unpark(ticket)
		expectedError := parking.InvalidTicketError{}

		assert.Exactly(t, expectedError, err2)
		assert.ErrorContains(t, err2, expectedError.Error())
		assert.Nil(t, returnedCar2)
	})
}