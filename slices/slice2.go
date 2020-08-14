package main

import "fmt"

func main() {
	//Crea un slice de tipo string
	//a diferencia de los array, los slice no llevan tamaño
	//predefinido
	frutas := []string{"manzana", "pera"}
	fmt.Println(frutas)

	//Esta 2da forma de crear el slice indica que se va
	//a crear un slice de tipo string con capacidad 2
	frutas2 := make([]string, 2)
	frutas2[0] = "manzana"
	frutas2[1] = "pera"
	fmt.Println(frutas2)

	//Ahora si declaramos la longitud y la capacidad
	//ocurre que no podemos asignar un valor que esté fuera
	//de la longitud del slice.
	//En esta forma el segundo argumento es la longitud
	//el 3er argumento es la capacidad
	frutas3 := make([]string, 0, 10)
	// frutas3[5] = "mango" //error
	fmt.Println(frutas3)

	//Append hace una copia del slice y le agrega
	//un nuevo valor y devuelve un nuevo slice
	frutas4 := make([]string, 0, 10)
	frutas4 = append(frutas4, "mango")
	fmt.Println(frutas4)

	//Va a imprimir mango pero con 5 espacios antes que él
	//Esto sucede porque de esta forma ya se ha inicializado
	//el slice
	frutas5 := make([]string, 0, 10)
	frutas5 = frutas5[0:6]
	frutas5[5] = "mango"
	fmt.Println(frutas5)

	//Si colocamos este tipo de dato, se visualizará de
	//mejor manera porque imprime su valor inicial
	edades := make([]int, 0, 10)
	edades = edades[0:6]
	edades[5] = 8
	fmt.Println(edades)

	//para imprimir longitud de slice len
	//para imprimir capacidad de slice cap
	fmt.Println(len(edades))
	fmt.Println(cap(edades))

	edades2 := make([]int, 0, 3)
	capacidad := cap(edades2)

	fmt.Printf("La capacidad inicial es %v\n", capacidad)

	for i := 0; i < 25; i++ {
		edades2 = append(edades2, i)
		fmt.Println(edades2)
		fmt.Println(cap(edades2))
		if cap(edades2) != capacidad {
			capacidad = cap(edades2)
			fmt.Printf("La nueva capacidad es %v\n", capacidad)
		}
	}

	//Otra forma de declarar
	var edades3 = []int{1, 2}
	fmt.Println(edades3)

	//Otra forma
	var edades4 []int
	fmt.Println(edades4)
}
