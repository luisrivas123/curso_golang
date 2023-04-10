package main

/*
   Hazlo solo y únicamente cuando la librería externa
   que estés usando lo pida explícitamente
*/
import (
	"fmt"
	"log"
	"strconv"

	// Importar una librería sin usarla
	_ "math"
)

// Función para tomar los errores (ahorra mucho código)
func isError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	a := 2
	fmt.Printf("%T", a)

	fmt.Println("Hola mundo")

	// Ejemplo de uso dunción error
	_, err := strconv.Atoi("53a")
	isError(err)

	// Slice de interfaces (Úsalo con sabiduría)
	// Permite guardar diferentes tipos de datos en un mismo slice
	myList := []interface{}{"Hola", 12, 4.90}

	// Iterar sobre los distintos tipos de datos de ese slice
	for _, v := range myList {
		switch v.(type) {
		case int:
			fmt.Println("Es int")
		case string:
			fmt.Println("Es string")
		case float64:
			fmt.Println("Es float64")
		}
	}

	m := make(map[string]int)

	m["hola"] = 1

	// Nota, usalmente se usa "ok" para recibir la segunda variable
	value, ok := m["hello"]

	fmt.Println(value, ok)
	/*
	   Si existe, ok será "true"
	   Si no existe, ok será "false"

	   En este caso, ok es "false" porque no existe.
	*/

	// Punteros
	a := 10 // Variable int
	b := &a // "b" es el puntero de "a"
	c := *b // "c" adquiere el valor del puntero de "b", es decir toma el mismo valor de "a"

	// Comandos de Go modules
// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod

}