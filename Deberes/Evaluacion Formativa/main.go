package main

import (
	"errors"
	"fmt"
	"time"
)

type Medicina struct {
	nombre           string
	fabricante       string
	fechaFabricacion time.Time
	vidaUtilMeses    int
}

func NuevaMedicina(nombre, fabricante string, fecha time.Time, vida int) (Medicina, error) {
	if nombre == "" || fabricante == "" {
		return Medicina{}, errors.New("no pueden haber estar vacíos")
	}
	if vida <= 0 {
		return Medicina{}, errors.New("la vida útil debe ser mayor a 0")
	}
	return Medicina{
		nombre:           nombre,
		fabricante:       fabricante,
		fechaFabricacion: fecha,
		vidaUtilMeses:    vida,
	}, nil
}

func (m *Medicina) GetNombre() string { return m.nombre }
func (m *Medicina) SetNombre(n string) error {
	if n == "" {
		return errors.New("el nombre solo puede ser texto")
	}
	m.nombre = n
	return nil
}

func (m *Medicina) GetFabricante() string { return m.fabricante }
func (m *Medicina) CalcularCaducidad() time.Time {
	return m.fechaFabricacion.AddDate(0, m.vidaUtilMeses, 0)
}

type Tableta struct {
	Medicina
	dosisxtableta int
	Prescripcion  bool
}

type Jarabe struct {
	Medicina
	volumen int
	sabor   string
	dosisml float64
}

func NuevaTableta(n, f string, fecha time.Time, vida, dosis int, presc bool) (*Tableta, error) {
	if dosis <= 0 {
		return nil, errors.New("No hay dosis de Pastillas o Tabletas")
	}
	base, err := NuevaMedicina(n, f, fecha, vida)
	if err != nil {
		return nil, err
	}
	return &Tableta{Medicina: base, dosisxtableta: dosis, Prescripcion: presc}, nil
}

func NuevoJarabe(n, f string, fecha time.Time, vida, vol int, sabor string) (*Jarabe, error) {
	if vol <= 0 {
		return nil, errors.New("el volumen debe ser mayor a 0")
	}
	base, err := NuevaMedicina(n, f, fecha, vida)
	if err != nil {
		return nil, err
	}
	return &Jarabe{Medicina: base, volumen: vol, sabor: sabor}, nil
}

func (t Tableta) MostrarInfoTa() {
	fmt.Println("-------Tabletas-----")
	fmt.Println("Nombre", t.nombre)
	fmt.Println("Fabricante", t.fabricante)
	fmt.Println("Fechtab", t.fechaFabricacion)
	fmt.Println("Vida Util", t.vidaUtilMeses)
	fmt.Println("Dosis", t.dosisxtableta)
	fmt.Println("Prescripcion?", t.Prescripcion)

}

func (s Jarabe) MostrarInfoJa() {
	fmt.Println("-------Jarabe----")
	fmt.Println("Nombre", s.nombre)
	fmt.Println("Fabricante", s.fabricante)
	fmt.Println("Caducidad", s.vidaUtilMeses)
	fmt.Println("Fecha", s.fechaFabricacion)
	fmt.Println("Dosis", s.dosisml)
	fmt.Println("Sabor", s.sabor)

}

func main() {
	inventarioTablas := []*Tableta{}
	invertarioJabe := []*Jarabe{}

}
