package main

import "fmt"

func switchCase() {

	// Con condición específica cuando se conocen los resultados de la variable
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// sin condición
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es mmenor a 0")
	default:
		fmt.Println("No es una condición")

	}
}
