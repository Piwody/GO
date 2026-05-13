package main

import "fmt"

type Producto struct {
	Nombre        string
	Precio        float64
	CantidadStock int // Asegúrate de usar este nombre siempre
}

func calcularValorTotal(p Producto) float64 {
	return p.Precio * float64(p.CantidadStock)
}

func estaDisponible(p Producto) string {
	if p.CantidadStock > 0 {
		return "Disponible"
	}
	return "Sin stock"
}

func mostrarInfo(p Producto) {
	total := calcularValorTotal(p)
	estado := estaDisponible(p)

	fmt.Printf("Nombre: %-12s, Stock: %-5d, Precio: %-7.2f, Total: %-8.2f, Estado: %s\n",
		p.Nombre, p.CantidadStock, p.Precio, total, estado)
}

func main() {
	
	p1 := Producto{Nombre: "Camiseta", Precio: 19.99, CantidadStock: 10}
	p2 := Producto{Nombre: "Pantalon", Precio: 25.00, CantidadStock: 0}
	p3 := Producto{Nombre: "Medias Rojas", Precio: 0.99, CantidadStock: 12}

	inventario := []Producto{p1, p2, p3}

	fmt.Println("--- INVENTARIO DE PRODUCTOS ---")

	for _, prod := range inventario {
		mostrarInfo(prod)
	}
}