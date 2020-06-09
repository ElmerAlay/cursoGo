package main

import (
	"fmt"
	"strconv"
)

func main() {
	//--------------------------------------------------------------------------
	//------------------------------------ Notas -------------------------------
	//--------------------------------------------------------------------------

	// no se puede importar una librería que no se utiliza, no compilará
	// no se puede concatenar variables de distinto tipo, siempre deben ser iguales

	// para convertir un string a entero importamos strconv y utilizamos su funcion Atoi
	// para convertir un string a entero es necesario declarar otra variable que sería su error, esto lo representamos con "_"
	// ahora ya sería un entero

	edad := "22"
	edadint, _ := strconv.Atoi(edad)

	fmt.Println(edadint + 10)

	// para convertir un entero a string no es necesario colocar el "_"
	// también se puede hacer directamente en el print

	edad1 := 22
	edadstr := strconv.Itoa(edad1)

	fmt.Println("mi edad es " + edadstr)
	fmt.Println("mi edad es " + strconv.Itoa(edad1))
}
