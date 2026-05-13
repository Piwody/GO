package main

import "fmt"

func mostrarnumeros(numeros ...int) {
	fmt.Println(numeros)

}

// el mensaeje es varibale, y siempre tine que ir al final.

func saludar(mensaje string, nombre ...string) {
	fmt.Println(mensaje, nombre)
}

func sumarvariables(numeros ...int) int {
	suma := 0
	for _, n := range numeros {
		suma += n

	}
	return suma

}
func main() {
	saludar("Hola como estan?", "Mateo", "Juan", "Pedro")
	resultado := sumarvariables(1, 2, 3, 4, 5)
	res1 := sumarvariables(1, 2, 12, 4, 5)
	res2 := sumarvariables(1, 14, 12, 4, 5)

	objetos := make([]string, 2)

	objetos[0] = "teclado"

	objetos[1] = "mouse"

	objetos[3] = "Cámara"

	fmt.Println(" Inventario inical:", objetos)

	objetos = append(objetos, "Monitor")

	fmt.Println("Los agregados son:", objetos)

	fmt.Println("L suma de esto es:", resultado)
	fmt.Println("L suma de esto es:", res1)
	fmt.Println("L suma de esto es:", res2)
}
