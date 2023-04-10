package main

import "fmt"

// Clases: atributos y métodos
// Type = Tipo de dato
type car struct {
	// Atributos
	brand string
	year  int
}

func structClass() {
	// Instanciar un struct
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
