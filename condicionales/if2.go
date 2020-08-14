package main

import "fmt"

func main() {

	//Se pueden declarar variables dentro del if
	if edad := 17; edad > 18 {
		fmt.Println("Mayor de edad")
		fmt.Println(edad)
	} else {
		fmt.Println("Menor de edad")

		//La variable también se puede usar dentro del else
		fmt.Println(edad)
	}
}
