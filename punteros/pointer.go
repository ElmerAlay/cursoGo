package main

import "fmt"

func main() {

	//Es una dir de mem
	//En lugar del valor, tenemos la dir de mem dónde se encuentra el valor
	//x,y = as123d = 5
	//x = as123d = 6
	//y ? = 6
	// *t = *int, *string, *struct
	//valor zero = nil

	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*x)
	fmt.Println(*y)

	*x = 6
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*x)
	fmt.Println(*y)
}
