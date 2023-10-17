package car_test

import (
	"testing"

	"github.com/cristianino/design-patterns-golang-builder/pkg/builder"
	"github.com/stretchr/testify/assert"
)

func TestSportsCarBuilder(t *testing.T) {
	// Crear un constructor de carro deportivo
	builderSport := &builder.SportsCarBuilder{}

	// Crear un director con el constructor
	director := builder.NewCarBuildDirector(builderSport)

	// Construir el carro
	sportsCar := director.Construct()

	// Verificar que el carro deportivo se construyó correctamente
	assert.Equal(t, "Motor deportivo", sportsCar.Engine)
	assert.Equal(t, "Ruedas deportivas", sportsCar.Wheels)
	assert.Equal(t, "Chasis deportivo", sportsCar.Chassis)
}

func TestBuggyCarBuilder(t *testing.T) {
	// Crear un constructor de carro deportivo
	builderBuggy := &builder.BuggiesCarBuilder{}

	// Crear un director con el constructor
	director := builder.NewCarBuildDirector(builderBuggy)

	// Construir el carro
	buggyCar := director.Construct()

	// Verificar que el carro deportivo se construyó correctamente
	assert.Equal(t, "Motor de buggy", buggyCar.Engine)
	assert.Equal(t, "Ruedas de buggy", buggyCar.Wheels)
	assert.Equal(t, "Chasis de buggy", buggyCar.Chassis)
}
