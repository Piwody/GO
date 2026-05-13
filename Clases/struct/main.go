package main

import "fmt"

type Estudiante struct {
	Nombre string
	Nota   float64
}

//metodo
func suirpunto(e Estudiante) {
	e.Nota += 1

}
func subir2punto(e *Estudiante) {
	e.Nota += 1

}
func verinota(e Estudiante) string {
	if e.Nota >= 70 {
		return "Arueba"
	}
	return "Reprueba"

}

func main() {
	est := Estudiante{"Ana", 98.5}
	fmt.Println("Bienvenido:", est.Nombre)
	fmt.Println("Tu nota es", est.Nota)

	fmt.Println("Con esta calificacon:", est.Nota, "tu estas", verinota(est))
	suirpunto(est)

}
