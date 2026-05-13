package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// bufio porque Scanln se detiene en el primer espacio
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingresa una frase:")
	reader.Scan()
	texto := reader.Text()

	
	palabras := strings.Fields(strings.ToLower(texto))

	//map
	recuento := make(map[string]int)

	for _, p := range palabras {
		recuento[p]++
	}


	fmt.Println("\nConteo de palabras")
	for palabra, cantidad := range recuento {
		fmt.Printf("%s: %d veces\n", palabra, cantidad)
	}


	palabraTop, vecesTop := palabraMasFrecuente(recuento)

	fmt.Printf("La palabra que más se repite es '%s' %d veces.\n", palabraTop, vecesTop)
}


func palabraMasFrecuente(lista map[string]int) (string, int) {
	mayor := 0
	ganadora := ""

	for palabra, cantidad := range lista {
		if cantidad > mayor {
			mayor = cantidad
			ganadora = palabra
		}
	}
	return ganadora, mayor
}

