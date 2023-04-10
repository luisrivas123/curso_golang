package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func goRoutines() {
	// acumula un grupo de Goroutines y la libera
	var wg sync.WaitGroup

	fmt.Println("Hello")
	// Agregacmos la Gorutine
	wg.Add(1)

	go say("world", &wg)

	// Espere a que todas las Goroutines de WaiteGroup finalicen
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")
}
