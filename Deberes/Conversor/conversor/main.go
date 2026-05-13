package main

import (
	"fmt"
)

func main() {

	var Opcionmoneda string
	var dolares float64

	fmt.Println("Bienvenido elige que cambio quieres hacer")
	fmt.Println("1. Yuan")
	fmt.Println("2. Euro")
	fmt.Println("3. Libra Esterlina")
	fmt.Print("Opcion:...")

	fmt.Scan(&Opcionmoneda)

	switch Opcionmoneda {
	case "1":
		fmt.Println("Ingresa la cantida en dolares para Yuan: ")
		fmt.Scan(&dolares)
		fmt.Printf("El resaultado es %.2f Yuanes.", dolares*7.24)
	case "2":
		fmt.Println("Ingresa la cantida en dolares para Euro: ")
		fmt.Scan(&dolares)
		fmt.Printf("El resaultado es %.2f Euro.", dolares*0.92)
	case "3":
		fmt.Println("Ingresa la cantida en dolares para Libra Esterlina: ")
		fmt.Scan(&dolares)
		fmt.Printf("El resaultado es %.2f Libra esterlina.", dolares*0.79)
	default:
		fmt.Print("ERROR!! \nOpcion no valida, eligue bien sopa!!")
	}

}
