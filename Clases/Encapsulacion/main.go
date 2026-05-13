package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func (p Persona) GetNombre() string {
	return p.nombre
}

func (p Persona) GetEdad() int {
	return p.edad
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) SetEdad(edad int) {
	if edad >= 0 {
		p.edad = edad
	}

}

func (p Persona) MostrarInfo() {
	fmt.Println("Nombre:", p.GetNombre())
	fmt.Println("Edad:", p.GetEdad())
}

func main() {
	perso := Persona{}

	fmt.Println("**Bienvenido a encapsulacion******")
	perso.SetNombre("Mateo Cantos")
	perso.SetEdad(20)

	perso.MostrarInfo()
}
