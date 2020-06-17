package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, 4)

	fmt.Println(slice)
	fmt.Println(copia)

	//Copia el mínimo de elementos en ambos arreglos

	copy(copia, slice)
	fmt.Println(slice)
	fmt.Println(copia)

	slice2 := []int{5, 6, 7, 8}
	copia2 := make([]int, 0)
	fmt.Println(slice2)
	fmt.Println(copia2)

	copy(copia2, slice2)
	fmt.Println(slice2)
	fmt.Println(copia2)

	slice3 := []int{1, 2, 3, 4}
	copia3 := make([]int, 2)
	fmt.Println(slice3)
	fmt.Println(copia3)

	copy(copia3, slice3)
	fmt.Println(slice3)
	fmt.Println(copia3)

	slice4 := []int{5, 6, 7, 8}
	copia4 := make([]int, len(slice4))
	fmt.Println(slice4)
	fmt.Println(copia4)

	copy(copia4, slice4)
	fmt.Println(slice4)
	fmt.Println(copia4)
}
