package main

import (
	"fmt"
)

func main() {
	//Todos los valores son del mismo tipo
	//Toman el valor por defecto del tipo de dato
	var arreglo [10]int
	fmt.Println(arreglo)

	//Se inicializan a los valores que queremos
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//Puede que haya menos valores al tamaño del arreglo, los
	//que faltan se les da el valor por defecto
	arr3 := [3]int{1, 2}
	fmt.Println(arr3)

	//Para obtener una posicion es [pos]
	//La pos inicia en 0 y terminal en tam-1
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	//Pueden existir arreglos con n dimensiones
	//Para acceder a un valor es [x][y]
	var matriz [2][2]int
	fmt.Println(matriz)
}
