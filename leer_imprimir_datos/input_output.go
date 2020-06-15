package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	edad := 26
	fmt.Printf("I am %d years old.\n", edad)

	name := "Elmer"
	fmt.Printf("Hi, my name is %v and I am %v years old.\n", name, edad)

	value := 26.54
	fmt.Printf("%e - %b - %f\n", value, value, value)

	var age int
	var name2 string
	fmt.Print("What is your age? ")
	fmt.Scanf("%d\n", &age)
	fmt.Print("What is your name? ")
	fmt.Scanf("%v\n", &name2)
	fmt.Printf("Your name is %s and you have %v years old.\n", name2, age)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name3, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Hi " + name3)
	}

	// ---------------------------------------------------------------
	// -------------------------- Notas ------------------------------
	// ---------------------------------------------------------------
	// %d tomará un valor numérico
	// %v tomará el valor por defecto del dato
	// %t tomará un valor booleano
	// %b %e %f para valores flotantes
	// %s para cadenas
	// Printf, Print, Println para imprimir datos
	// %v en Scanf sólo leerá cadenas
}
