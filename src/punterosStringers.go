package main

import (
	pk "curso_golang/src/mypackage"
	"fmt"
)

func punterosStringers() {

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	var myPc pk.Pc

	myPc.Ram = 16
	myPc.Disk = 200
	myPc.Brand = "msi"

	fmt.Println(myPc)

	myPc.Ping()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)

}
