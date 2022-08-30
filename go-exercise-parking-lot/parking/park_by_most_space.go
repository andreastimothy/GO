package parking

type ParkByMostSpace struct {}

func (style *ParkByMostSpace) SelectLot(availableLot []*Lot) *Lot {
	maxSpace := 0
	index := 0
	for i := range availableLot {
		availableFreeSpace := availableLot[i].Capacity - len(availableLot[i].Slot)
		if availableFreeSpace > maxSpace {
			maxSpace = availableFreeSpace
			index = i
		}
	}
	return availableLot[index]
}