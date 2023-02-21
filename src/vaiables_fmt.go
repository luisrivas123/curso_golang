package main

import (
	"fmt"
)

func main_variables() {
	// fmt.Println("Hola mundo")

	// Declaración de constantes
	// const pi float64 = 3.14
	// const pi2 = 3.1416

	// fmt.Println("pi:", pi)
	// fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	// base := 12
	// var altura int = 14
	// var area int

	// area = base * (altura / 2)
	// fmt.Println("Area:", area)

	// Zero values
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a, b, c, d)

	// Clase fmt

	helloMessage := "Hello"
	worldMessage := "world"

	// Println

	fmt.Println(helloMessage, worldMessage)

	// Printf

	nombre := "Luis"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
