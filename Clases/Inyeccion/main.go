package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}
type empleado struct {
	Persona
	cargo   string
	salario float64
}

func (e empleado) mostrarInfo() {
	fmt.Println(e.nombre, e.cargo, e.salario, e.edad)

}

func main() {
	empleado2 := empleado{
		Persona: Persona{"Nicolas", 20},
		cargo:   "Ingeniero de datos",
		salario: 1200,
	}
	empleado2.mostrarInfo()
}
