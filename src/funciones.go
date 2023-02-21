package main

import (
	"fmt"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnFuction(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func funciones() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")
	value := returnFuction(2)
	fmt.Println("Value:", value)
	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)
	value1P, _ := doubleReturn(2)
	fmt.Println("value1:", value1P)
}
