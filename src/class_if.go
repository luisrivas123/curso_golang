package main

import (
	"fmt"
)

func login(user string, password int) {
	usuario := "Luis"
	pass := 123

	acces := usuario == user && pass == password
	fmt.Println(acces)

}

func class_if() {
	valor := 40

	if valor%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("No es par")
	}

	login("Lui", 123)

	//

	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	if isValidUser("Alpha", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}

}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}
