package main

import "fmt"

type Producto struct {
	iD       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) *Producto {
	return &Producto{
		iD:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

////
func (p *Producto) GetID() int         { return p.iD }
func (p *Producto) GetNombre() string  { return p.nombre }
func (p *Producto) GetPrecio() float64 { return p.precio }
func (p *Producto) GetCantidad() int   { return p.cantidad }

func (p *Producto) SetPrecio(nuevoPrecio float64) {
	p.precio = nuevoPrecio
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	p.cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Producto: %-10s | $: %.2f | Stock: %d\n",
		p.iD, p.nombre, p.precio, p.cantidad)
}

/////

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() *Inventario {
	return &Inventario{
		productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {

	inv.productos[p.GetID()] = p
}

func (inv *Inventario) EliminarProducto(id int) {
	delete(inv.productos, id)
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv *Inventario) ListarProductos() {
	fmt.Println("\n--- Inventario ---")
	if len(inv.productos) == 0 {
		fmt.Println("Inventario vacío.")
		return
	}
	for _, p := range inv.productos {
		p.MostrarInfo()
	}
}

/////
func main() {

	tienda := NuevoInventario()

	p1 := NuevoProducto(1, "Papaya", 1.20, 5)
	p2 := NuevoProducto(2, "Manzana", 0.50, 20)
	p3 := NuevoProducto(3, "Aguacatae", 1.00, 12)

	tienda.AgregarProducto(*p1)
	tienda.AgregarProducto(*p2)
	tienda.AgregarProducto(*p3)

	tienda.ListarProductos()

	p1.SetPrecio(1.99)
	p1.SetCantidad(8)
	tienda.AgregarProducto(*p1)

	fmt.Println("\nActualizando...")
	tienda.ListarProductos()

	tienda.EliminarProducto(3)
	fmt.Println("\nAguacate, eliminado:")
	tienda.ListarProductos()
}
