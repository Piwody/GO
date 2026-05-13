package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bienvenido a mi programa de ciclos con GO")
	/*
	   	//for con incremento
	   	fmt.Println(" Bucle Normal ")
	   	for i := 0; i <= 10; i++ {

	   		fmt.Println(i)
	   	}

	   	// for infinito
	   	for {
	   		fmt.Println("error")
	   		break
	   	}


	     // for range
	     println("For in range")

	     for rango :=  range [6] int{} {
	       fmt.Println(rango)
	     }
	
	//eje

	println("Numeros hasta el 50")

	for num := range [10]int{} {
		if num%2 == 0 {
			continue
		} else {
			fmt.Println(num)
		}

	}

	// ejerc
	fmt.Println("Ingresa un numero")
	var num int
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")

	}
*/

	//ejer 2

	fmt.Println("Ingresa tu nombre")
	scanner := bufio.NewReader(os.Stdin)
	palabras, _ := scanner.ReadString('\n')
	fmt.Println("Tu frase es:", palabras)

}
