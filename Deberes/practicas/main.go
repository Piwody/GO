package main

import (
	"fmt"
)

/*
var nombre string
var edad int
var peso float64
var Fuma bool

func main() {

	fmt.Println("Bienvenido, ingresa tu nombre completo:")
	fmt.Scanln(&nombre)

	fmt.Println("Ingresa tu edad")
	fmt.Scanln(&edad)

	fmt.Println("Ingresa tu peso en kg")
	fmt.Scanln(&peso)

	fmt.Println("Hola", nombre, "Bienvenido a esta muestra, tu peso es de:", peso)
}


func main() {
	saldo := 250.0
	var retiro float64

	fmt.Println("Bienvenido al Cajero UIDE")
	fmt.Println("Su saldo actual es: $", saldo)

	fmt.Println("Ingresa cuanto desea retirar......")
	fmt.Scan(&retiro)

	if retiro <= 0 {
		fmt.Println("No se puede retirar ese monto, ingrese otro monto.....")
	} else if retiro > saldo {
		fmt.Println("Saldo insuficiente, ingrese otro monto.....")
	} else {
		saldo = saldo - retiro
		fmt.Println("Retiro exitoso.....")
		fmt.Println("Su saldo es de:", saldo)
	}
}
*/

func main() {
	juego := make([]string, 3)

	juego[0] = "Fornite"
	juego[1] = "CS2"
	juego[2] = "Valo"

	fmt.Println("Juegos disponibles\n", juego)

}
