package builder

import "github.com/cristianino/design-patterns-golang-builder/internal/car"

// Constructor Concreto para Carro Buggy
type BuggiesCarBuilder struct {
	car car.Car
}

func (b *BuggiesCarBuilder) BuildEngine() {
	b.car.Engine = "Motor de buggy"
}

func (b *BuggiesCarBuilder) BuildWheels() {
	b.car.Wheels = "Ruedas de buggy"
}

func (b *BuggiesCarBuilder) BuildChassis() {
	b.car.Chassis = "Chasis de buggy"
}

func (b *BuggiesCarBuilder) GetCar() car.Car {
	return b.car
}
