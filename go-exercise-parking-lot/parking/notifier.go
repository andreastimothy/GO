package parking

type LotFullNotifier interface {
	NotifyLotFull(*Lot)
}

type LotAvailableNotifier interface {
	NotifyLotAvailable(*Lot)
}