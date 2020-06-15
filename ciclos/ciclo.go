package main

import (
	"fmt"
)

func main() {
	//No es obligatorio que vengan todos los datos

	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}

	e := 0
	for e < 5 {
		fmt.Println(e + 1)
		e++
	}

	f := 0
	for {
		if f == 2 {
			f++
			continue
		}
		fmt.Println(f)

		f++ //debe ir aquí porque sino entra a un bucle infinito

		if f > 5 {
			break
		}
	}
}
