package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nombre string
	nombre = "Elmer"
	fmt.Println(nombre)

	nombre = nombre + " Alay"
	fmt.Println(nombre)

	//Error, no realiza el salto de linea
	/*
		nombre = "Elmer
		Alay"
		fmt.Println(nombre)
	*/

	//No realiza el salto de linea
	nombre = `nombre\nAlay`
	fmt.Println(nombre)

	//Sí realiza el salto de línea
	nombre = `Elmer
Alay`
	fmt.Println(nombre)

	//Realiza el salto de linea
	nombre = "Elmer\nAlay"
	fmt.Println(nombre)

	//tamaño
	fmt.Println(len(nombre))

	//Podemos acceder a una posicion
	//pero lo devuelve en código asccii
	fmt.Println(nombre[0])

	//imprime el tipo de dato
	fmt.Println(reflect.TypeOf(nombre[0]))

	//imprime caracter
	fmt.Println(string(nombre[0]))

	for pos, letter := range nombre {
		fmt.Printf("%d, \"%s\"\n", pos, string(letter))
	}
}
