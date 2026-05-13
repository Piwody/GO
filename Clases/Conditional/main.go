package main

import "fmt"

func main() {
	flag := false
	var user string
	privileges := false

	if flag {
		user = "admin"
	} else {
		user = "guest"
	}

	edad := 0
	fmt.Scan(&edad)

	if edad >= 18 && user == "admin" {
		fmt.Println("Bienvenido, tienes acceso completo al sistema")
		privileges = true
		fmt.Println(privileges)
	} else if edad < 18 && user == "guest" {
		fmt.Println("Acceso denegado")

	} else {
		fmt.Println("No tienes edad suficiente")
		privileges = false
		fmt.Println(privileges)

	}

	if n:=9 ; n < 10 {
		fmt.Println(n ,"es menor que 10")


		

	}

}

