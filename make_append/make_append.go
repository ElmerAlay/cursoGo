package main

import "fmt"

func main() {
	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	//Capacidad
	fmt.Println(cap(slice))
	slice2 := make([]string, 2)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := make([]int, 0, 2)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 2)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 2)
	slice3 = append(slice3, 3)
	slice3 = append(slice3, 4)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	//La capacidad es distinto de la longitud hasta el momento
	//en que la longitud se iguala, por ejemplo,
	//tengo un slice con capacidad 3, pero sólo 2 elementos
	//la longitud es 2 y la capacidad es 3
	//en el momento en que agrego otros 2 valores ahora
	//la longitud es 4 y la capacidad 4
}
