package main

import "fmt"

func main() {
	//Son como diccionarios
	//Se almacenan claves y valores

	//Significa que creamos un mapa
	//con llaves de tipo string
	//valores de tipo entero
	frutas := make(map[string]int)

	frutas["mango"] = 5
	frutas["manzana"] = 10

	fmt.Println(frutas)

	//De esta forma ocupa menos espacio. por lo tanto,
	//es más eficiente
	frutas2 := make(map[string]int, 2)
	frutas2["mango"] = 5
	frutas2["manzana"] = 10
	fmt.Println(frutas2)

	//Elimina la clave y el valor contenido en mango
	delete(frutas, "mango")
	fmt.Println(frutas)

	//Se puede eliminar una clave que ni siquiera existe
	delete(frutas, "banana")
	fmt.Println(frutas)

	//Para recorrer el mapa
	for key, value := range frutas2 {
		fmt.Printf("%s: %d\n", key, value)
	}
}
