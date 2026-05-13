package main

import (
	"fmt"
)

type CuentaBancaria struct {
	titular string
	saldo   float64
}

func (c *CuentaBancaria) setTitular(titualar string) {
	c.titular = titualar
}

func (c *CuentaBancaria) setSaldo(Saldo float64) {
	c.saldo = Saldo
}

func (c *CuentaBancaria) getTitular() string {
	return c.titular
}

func (c *CuentaBancaria) getSaldo() float64 {
	return c.saldo
}

func NewCuenta(titular string, saldo float64) *CuentaBancaria {
	return &CuentaBancaria{titular, saldo}

}

func (c CuentaBancaria) NuevaCuenta(titular string, saldo float64) {
	c.setTitular(titular)
	c.setSaldo(saldo)
}

func (c *CuentaBancaria) Depositar(monto float64) {
	c.saldo += monto
}

func (c *CuentaBancaria) Retiro(monto float64) {
	if monto > c.getSaldo() {
		fmt.Println("No hay fondos suficeitnes")
	} else {
		c.saldo -= monto
		fmt.Println("Retiro exitoso. El retiro ha sido de:", monto)
	}

}

func (c *CuentaBancaria) mostrarInfo() {
	fmt.Println("Titular:", c.getTitular())
	fmt.Println("Saldo:", c.getSaldo())
}

func (c *CuentaBancaria) MostrarSaldo() {
	fmt.Println("Saldo:", c.saldo)
}

func main() {

	fmt.Println("-------Bienvenido-------")
	fmt.Println("-------Datos de tu cuenta-------")
	cuenta := NewCuenta("Mateo", 1000)
	cuenta.mostrarInfo()
	cuenta.Depositar(500)
	cuenta.Retiro(1)
	cuenta.MostrarSaldo()

}
