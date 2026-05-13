package main

import (
	"fmt"
	"slices"
)

func main() {
	slicescadena := make([]string, 3)

	slicescadena[0] = "a"
	slicescadena[1] = "b"
	slicescadena[1] = "b"

	fmt.Println("Datos:", slicescadena)
	slicescadena = append(slicescadena, "1234")
	fmt.Println("Datos modificados:", slicescadena)

	slice := []string{"g", "h", "k"}
	slice2 := []string{"g", "h", "i"}

	if slices.Equal(slice, slice2) {
		fmt.Println("Los dos slicess sn iguales")
	} else {
		fmt.Println("Los dos slicess son diferentes")

	}
}
