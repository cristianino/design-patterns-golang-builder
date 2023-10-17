package builder

import "github.com/cristianino/design-patterns-golang-builder/internal/car"

// Constructor Concreto para Carro Deportivo
type SportsCarBuilder struct {
	car car.Car
}

func (b *SportsCarBuilder) BuildEngine() {
	b.car.Engine = "Motor deportivo"
}

func (b *SportsCarBuilder) BuildWheels() {
	b.car.Wheels = "Ruedas deportivas"
}

func (b *SportsCarBuilder) BuildChassis() {
	b.car.Chassis = "Chasis deportivo"
}

func (b *SportsCarBuilder) GetCar() car.Car {
	return b.car
}
