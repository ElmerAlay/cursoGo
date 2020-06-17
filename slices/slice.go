package main

import "fmt"

func main() {
	//Arreglo de tamaño dinámico
	matriz := []int{1, 2, 3, 4}
	fmt.Println(matriz)

	var matriz2 []int
	if matriz2 == nil {
		fmt.Println("Vacío")
	} else {
		fmt.Println(matriz2)
		fmt.Println(len(matriz2))
	}

	//puntero al arreglo
	//longitud del arreglo al que apunta
	//capacidad

	arreglo := [3]int{1, 2, 3}

	//Tomará los valores desde el inicio (pos 0) hasta la pos 2
	//del arreglo
	slice := arreglo[:2]
	fmt.Println(slice)
	//Tomará la primera posición hasta la posicion 3
	slice2 := arreglo[0:3]
	fmt.Println(slice2)
	//Tomará el inicio del arreglo hasta el final
	slice3 := arreglo[0:]
	fmt.Println(slice3)
	//siempre del lado izquierdo se empezará [0,1,2,3]
	//Siempre del lado derecho se tomará el valor normal [1,2,3]
}
