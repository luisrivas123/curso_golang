package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func channelsRangeSelect() {
	c := make(chan string, 2)
	c <- "mensaje1"
	c <- "mensaje2"

	fmt.Println(len(c), cap(c))

	// Range y close
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	// select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; 1 < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
