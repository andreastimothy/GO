package parking

type Car struct {
	Plate string
}

func NewCar(plate string) *Car {
	return &Car{Plate: plate}
}

func GetPlate(car *Car) string {
	return car.Plate
}