package main

import "fmt"

func main() {
	//--------------------------------------------------------------------------
	//------------------------------------ Notas -------------------------------
	//--------------------------------------------------------------------------

	// 1. Todas las variables declaradas deben ser utilizadas, de lo contrario no dejará compilar el archivo

	// 2. La primera forma de declarar variables es de la siguiente forma
	//var nombre_var tipo_dato
	var a int //valor por defecto 0
	a = 2
	var b string //valor por defecto cadena vacía
	var c bool   //valor por defecto false

	// 3. La segunda forma es
	//en esta forma no es necesario colocar el tipo de dato
	//el tipo de la variable será igual al tipo del valor que le estamos asignando
	//ya no podrá cambiar su tipo luego
	d := 8

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
