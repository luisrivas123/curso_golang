package main

import (
	// pk "curso_golang/src/mypackage"
	"fmt"
)

func modificadoresAcceso() {
	// var myCar mypackage.CarPublic
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	pk.PrintMessage("Hola Platzi")
}
