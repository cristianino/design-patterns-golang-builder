package builder

import "github.com/cristianino/design-patterns-golang-builder/internal/car"

// Interfaz Builder
type CarBuilder interface {
	BuildEngine()
	BuildWheels()
	BuildChassis()
	GetCar() car.Car
}

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

// Director
type CarBuildDirector struct {
	builder CarBuilder
}

func NewCarBuildDirector(builder CarBuilder) *CarBuildDirector {
	return &CarBuildDirector{builder: builder}
}

func (d *CarBuildDirector) Construct() car.Car {
	d.builder.BuildEngine()
	d.builder.BuildWheels()
	d.builder.BuildChassis()
	return d.builder.GetCar()
}
