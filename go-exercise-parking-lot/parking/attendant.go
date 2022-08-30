package parking

type Attendant struct {
	Lots []*Lot
	AvailableLots []*Lot
	Style LotSelector
}

func NewAttendant(lot []*Lot) *Attendant {
	availableLot := make([]*Lot, len(lot))
	copy(availableLot, lot)

	attendant := &Attendant{
		Lots: lot,
		AvailableLots: availableLot,
		Style: &ParkSequentially{},
	}

	for _, lot := range lot {
		lot.AddSubscriberFull(attendant)
		lot.AddSubscriberAva(attendant)
	}
	return attendant
}

func (att *Attendant) AttPark(car *Car) (*Ticket, error) {
	for _, slot := range att.Lots {
		if slot.IsCarExist(car) {
			return nil, CarAlreadyParkedError{}
		}
		if att.AvailableLots != nil {
			return att.Style.SelectLot(att.AvailableLots).Park(car)
		}
	}
	return nil, ParkingIsFullError{}
}

func (att *Attendant) AttUnpark(ticket *Ticket) (*Car, error) {
	for _, slot := range att.Lots {
		if slot.IsCarExist(slot.Slot[ticket]) {
			return slot.Unpark(ticket)
		}
	}
	return nil, InvalidTicketError{}
}

func (att *Attendant) FindIndex (lot *Lot) int {
	for index := range att.Lots {
		if att.AvailableLots[index] == lot {
			return index
		}
	}
	return -1
}

func (att *Attendant) NotifyLotFull(lot *Lot) {
	index := att.FindIndex(lot)

	att.AvailableLots = append(att.AvailableLots[:index], att.AvailableLots[index + 1:]...)
}

func (att *Attendant) NotifyLotAvailable(lot *Lot) {
	att.AvailableLots = append(att.AvailableLots, lot)
}

func (att *Attendant) NewStyle(style LotSelector) {
	att.Style = style
}