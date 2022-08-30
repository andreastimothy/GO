package parking

type ParkByMostCapacity struct {}

func (style *ParkByMostCapacity) SelectLot(availablelot []*Lot) *Lot {
	index := 0
	for i, max := range availablelot {
		if max.Capacity > availablelot[index].Capacity {
			availablelot[index].Capacity = max.Capacity
			index = i
		}
	}
	return availablelot[index]
}