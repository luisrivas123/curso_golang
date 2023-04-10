package mypackage

import "fmt"

// Comentar que hace
// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMesaage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessage1(text string) {
	fmt.Println(text)
}

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

// Método
func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "Pong")
}

func (myPc *Pc) DuplicateRam() {
	myPc.Ram = myPc.Ram * 2
}

func (myPc Pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.Ram, myPc.Disk, myPc.Brand)
}
