package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	var lowerText = strings.ToLower(text)
	fmt.Println(lowerText)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(lowerText[i])
	}

	if lowerText == textReverse {
		fmt.Println("Es un palíndromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}

func sliceRange() {
	slice := []string{"hola", "que", "hace"}

	// for i, valor := range slice {
	// 	fmt.Println(i, valor)
	// }
	// for _, valor := range slice {
	// 	fmt.Println(valor)
	// }
	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("AmoR A rOma")
}
