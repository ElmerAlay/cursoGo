package main

import "fmt"

func main() {
	//No se puede cambiar el tamaño
	var edades [3]int

	//Inicia en 0
	edades[0] = 18
	edades[1] = 25
	edades[2] = 32
	//edades[3] = 11 //error
	fmt.Println(edades)

	//Para recorrer
	for index, value := range edades {
		fmt.Printf("%d -> %d\n", index, value)
	}

	//otra forma de declarar
	edades2 := [3]int{1, 2, 3}
	fmt.Println(edades2)
}
