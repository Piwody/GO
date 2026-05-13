package main

import (
	"errors"
	"fmt"
)

func divi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("nO PUEDES DIVIDR PARA CERO")
	}
	return a / b, nil
}

func main() {
	res, err := divi(3, 0)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("El resultaso es:", res)
	}
}
