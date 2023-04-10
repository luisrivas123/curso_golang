package main

import "fmt"

func maps() {
	// Maps, Diccionarios

	m := make(map[string]int)

	m["jose"] = 14
	m["pedro"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["josep"]
	fmt.Println(value, ok)

}
