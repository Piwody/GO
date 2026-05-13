package main

import (
	"fmt"
)

// aarray es fijo [] con numero
//slyce es variable [] sin numeor

// arreglo nos sirve para guardar una coleccion de datos del mismo tipo

func main() {
	var arreglo [5]int
	fmt.Println("Arrelgo:", arreglo, "\nTamano:", len((arreglo)))
	arreglo[4] = 100
	fmt.Println("Arreglo:", arreglo, "\n En la posicion 4 esta:", arreglo[4])
	array2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Arreglo inicializado :", array2)
	array3 := [...]int{1, 2, 3, 3, 34, 5, 3}
	fmt.Println("Arreglo variatico :", array3, "\nTamano:", len((array3)))

}
