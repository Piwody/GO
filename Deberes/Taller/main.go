package main

import (
	"fmt"
)

func main() {

	materias := []string{"Programacion", "Ciberseguridad", "Redes", "Calculo"}

	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 10.0},
		{6.0, 7.0, 8.0, 6.5},
		{10.0, 9.5, 9.8, 9.2},
		{5.5, 6.0, 4.0, 7.0},
		{8.0, 8.0, 8.0, 8.0},
		{7.2, 6.8, 9.0, 8.5},
	}

	var sumaTotalClase float64
	promediosClase := make([]float64, 0)

	fmt.Println("Por Materia")

	for i := 0; i < len(notas); i++ {
		fmt.Printf("\nEstudiante %d:\n", i+1)

		notasEstudiante := notas[i][:] // para calculaar estaditicas al final 

		for j := 0; j < len(materias); j++ {
			fmt.Printf(" - %-15s: %.2f\n", materias[j], notasEstudiante[j])
		}

		sum, max, min := calcularEstadisticas(notasEstudiante) /// hacer funcion abajo 
		promedio := sum / float64(len(notasEstudiante))

		promediosClase = append(promediosClase, promedio)
		sumaTotalClase += promedio

		fmt.Printf(" RESUMEN: Promedio: %.2f | Máxima: %.2f | Mínima: %.2f\n", promedio, max, min)
		fmt.Println(".....................")
	}

	promedioGeneral := sumaTotalClase / float64(len(promediosClase))
	fmt.Printf("\nPromedio de clase : %.2f\n", promedioGeneral)
}

func calcularEstadisticas(notas []float64) (suma, max, min float64) {
	max = notas[0]
	min = notas[0]
	for _, nota := range notas {
		suma += nota
		if nota > max {
			max = nota
		}
		if nota < min {
			min = nota
		}
	}
	return
}
