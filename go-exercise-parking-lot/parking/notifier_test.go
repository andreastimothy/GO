package parking_test

import (
	"testing"
	"git.garena.com/sea-labs-id/batch-02/shared-projects/go-exercise-parking-lot/parking/mocks"
	"git.garena.com/sea-labs-id/batch-02/shared-projects/go-exercise-parking-lot/parking"
)

func TestNotifyLotFull(t *testing.T) {
	t.Run("should notify when parking is full", func(t *testing.T) {
		test := mocks.NewLotFullNotifier(t)
		car := parking.NewCar("B 1234 ABC")
		parkinglot := parking.NewParkingLot(1)
		
		test.On("NotifyLotFull", parkinglot).Return(true, nil)

		parkinglot.AddSubscriberFull(test)
		parkinglot.Park(car)

		test.AssertExpectations(t)
		test.AssertNumberOfCalls(t, "NotifyLotFull", 1)
	})
}

func TestNotifyLotAvailable(t *testing.T) {
	t.Run("should notify when parking is available", func(t *testing.T) {
		test := mocks.NewLotAvailableNotifier(t)
		car := parking.NewCar("B 1234 ABC")
		parkinglot := parking.NewParkingLot(1)
		ticket, _ := parkinglot.Park(car)

		test.On("NotifyLotAvailable", parkinglot).Return(true, nil)

		parkinglot.AddSubscriberAva(test)
		parkinglot.Unpark(ticket)
		
		test.AssertExpectations(t)
		test.AssertNumberOfCalls(t, "NotifyLotAvailable", 1)
	})
}