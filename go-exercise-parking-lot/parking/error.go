package parking

type ParkingIsFullError struct {}
type InvalidTicketError struct {}
type CarAlreadyParkedError struct {}

func (err ParkingIsFullError) Error() string {
	return "no available position"
}

func (err InvalidTicketError) Error() string {
	return "unrecognized parking ticket"
}

func (err CarAlreadyParkedError) Error() string {
	return "car already parked"
}