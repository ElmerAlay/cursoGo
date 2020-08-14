package main

import "fmt"

func main() {
	nombre := "Elmer"
	for index, value := range nombre {
		fmt.Printf("Tipos de datos: %T, %T\n", index, value)
		fmt.Printf("Valor en número: %d, %d\n", index, value)
		fmt.Printf("Valor en cadena: %d, %s\n", index, string(value))
	}

	//Si no queremos usar una variable reemplazamos por un _
	for _, value := range nombre {
		fmt.Printf("Tipos de datos: %T\n", value)
		fmt.Printf("Valor en número: %d\n", value)
		fmt.Printf("Valor en cadena: %s\n", string(value))
	}

	//Si no queremos usar ninguna variable
	//Notar que no usamos :=, sino que se usa =
	for _, _ = range nombre {
		fmt.Printf("a\n")
	}

	animales := [2]string{"perro", "gato"}
	for _, value := range animales {
		fmt.Println(value)
	}

	frutas := []string{"manzana", "pera", "mango"}
	for _, value := range frutas {
		fmt.Println(value)
	}

	//Para los mapas la primera variable es
	//la llave
	idiomas := map[string]string{"es": "Español", "en": "Inglés"}
	for llave, value := range idiomas {
		fmt.Printf("%s, %s\n", llave, value)
	}

	idiomas2 := map[int]string{1200: "Español", 1300: "Inglés"}
	for llave, value := range idiomas2 {
		fmt.Printf("%d, %s\n", llave, value)
	}
}
