package myapp

import (
	"fmt"

	"github.com/cristianino/design-patterns-golang-builder/pkg/builder"
)

func Run() {
	var carDirector *builder.CarBuildDirector

	// Crear un constructor concreto para un carro deportivo
	sportsCarBuilder := &builder.SportsCarBuilder{}

	// Crear un constructor concreto para un carro buggy
	buggyCarbuilder := &builder.BuggiesCarBuilder{}

	// Crear un director con el constructor
	carDirector = builder.NewCarBuildDirector(sportsCarBuilder)

	// Construir el carro
	sportsCar := carDirector.Construct()

	carDirector = builder.NewCarBuildDirector(buggyCarbuilder)

	buggyCar := carDirector.Construct()

	// Obtener y mostrar el resultado
	fmt.Printf("Carro Deportivo:\nMotor: %s\nRuedas: %s\nChasis: %s\n",
		sportsCar.Engine, sportsCar.Wheels, sportsCar.Chassis)

	// Obtener y mostrar el resultado
	fmt.Printf("Carro Buggy:\nMotor: %s\nRuedas: %s\nChasis: %s\n",
		buggyCar.Engine, buggyCar.Wheels, buggyCar.Chassis)
}
