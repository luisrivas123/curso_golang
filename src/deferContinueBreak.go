package main

import "fmt"

func deferContinueBreak() {
	// Defer Es lo último que se ejecuta antes de terminar la programa
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Es 8")
			break
		}

	}
}
