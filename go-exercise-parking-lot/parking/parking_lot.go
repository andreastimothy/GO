package parking

type Lot struct {
	Slot     map[*Ticket]*Car
	Capacity int
	SubscribersFull []LotFullNotifier
	SubscribersAva []LotAvailableNotifier
}

func NewParkingLot(capacity int) *Lot {
	return &Lot{
		Slot:     make(map[*Ticket]*Car),
		Capacity: capacity,
	}
}

func (lot *Lot) IsFull() bool {
	return len(lot.Slot) == lot.Capacity
}

func (lot *Lot) IsCarExist(car *Car) bool {
	for _, value := range lot.Slot {
		if value.Plate == car.Plate {
			return true
		}
	}
	return false
}

func (lot *Lot) Park(car *Car) (*Ticket, error) {
	if lot.IsFull() {
		return nil, ParkingIsFullError{}
	}
	if lot.IsCarExist(car) {
		return nil, CarAlreadyParkedError{}
	}
	newTicket := NewTicket(car.Plate)
	lot.Slot[newTicket] = car

	if lot.IsFull() {
		lot.NotifySubscriberFull()
	}

	return newTicket, nil
}

func (lot *Lot) Unpark(ticket *Ticket) (*Car, error) {
	car, exist := lot.Slot[ticket]
	if !exist {
		return nil, InvalidTicketError{}
	}
	delete(lot.Slot, ticket)

	if len(lot.Slot) == lot.Capacity - 1 {
		lot.NotifySubscriberAva()
	}

	return car, nil
}

func (lot *Lot) AddSubscriberFull(subscriber LotFullNotifier) {
	lot.SubscribersFull = append(lot.SubscribersFull, subscriber)
}

func (lot *Lot) NotifySubscriberFull() {
	for _, subs := range lot.SubscribersFull {
		subs.NotifyLotFull(lot)
	}
}

func (lot *Lot) AddSubscriberAva(subscriber LotAvailableNotifier) {
	lot.SubscribersAva = append(lot.SubscribersAva, subscriber)
}

func (lot *Lot) NotifySubscriberAva() {
	for _, subs := range lot.SubscribersAva {
		subs.NotifyLotAvailable(lot)
	}
}