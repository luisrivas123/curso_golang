package main

import "fmt"

func arraySlice() {
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 8, 3, 4, 9, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[3:6])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice[7])

	// Append nuevaa lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
