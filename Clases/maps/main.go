package main

/*func main() {
	var n int

	fmt.Println("Ingresa el numero de estudiantes:")
	fmt.Scanln(&n)

	estuidiantes:= make(map [string]float64)
	var nombre string
	var nota float64

	for i:=0; i < n; i++{
		fmt.Println("Estudiante", i+1, ":")
		fmt.Scanln(&nombre)
		fmt.Println("Nota:")
		fmt.Scan(&nota)
		estuidiantes[nombre] = nota


	}
	fmt.Println("\n Datos Ingresados:")

	for name,grade := range estuidiantes{
		fmt.Println(name, ":", grade)


	}




}*/

func mejorestu(estudiante map[string]float64) (string, float64) {

	mayor := 0.0
	mejo := ""

	for nombre, nota := range estudiante {
		if nota > mayor {
			mayor = nota
			mejo = nombre
		}
	}
	return mejo, mayor

}

