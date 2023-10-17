package builder

import "github.com/cristianino/design-patterns-golang-builder/internal/car"

// Interfaz Builder
type CarBuilder interface {
	BuildEngine()
	BuildWheels()
	BuildChassis()
	GetCar() car.Car
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
