package parking

type ParkSequentially struct {}

func (style *ParkSequentially) SelectLot(availableLot []*Lot) *Lot {
	return availableLot[0]
}