package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Println(saludar())
	fmt.Println(saludar2("Elmer"))
	fmt.Println(saludar3("Elmer", "Alay"))
	fmt.Println(saludar4("Elmer", "Alay", 27))

	//Cuando una función devuelve 2 valores usamos
	//la siguiente estructura
	//recordando que si queremos omitir una usamos _
	saludo, mayor := saludar5("Elmer", "Alay", 27)
	fmt.Println(saludo)
	fmt.Println(mayor)

	saludo, mayor = saludar6("Elmer", "Alay", 27)
	fmt.Println(saludo)
	fmt.Println(mayor)

	saludo2, _, err := saludar7("Elmer", "Alay", 27)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(saludo2)

	slice("Elmer", "Edgardo", "alay", "yupe")

	//No se puede pasar un slice como parámetro
	nombres := []string{"Nicola", "Peltz"}
	//slice(nombres) //error

	//Pero sí se pueden pasar los valores del slice
	//Se hace de la siguiente manera
	slice(nombres...)
}

func saludar() string {
	return "Hola"
}

func saludar2(nombre string) string {
	return "Hola " + nombre
}

//Si 2 parámetros son del mismo tipo sólo se coloca una vez
//el tipo y las variables se separan por coma
func saludar3(nombre, apellido string) string {
	//Sprintf devuelve un string
	mensaje := fmt.Sprintf("Hola %s %s", nombre, apellido)
	return mensaje
}

//Si queremos unir un entero y un string sin necesidad de
//convertir podemos usar uint8
func saludar4(nombre, apellido string, edad uint8) string {
	mensaje := fmt.Sprintf("Hola %s %s, tienes %d años", nombre, apellido, edad)
	return mensaje
}

//Se pueden retornar 2 o más tipos de datos
//Lo importante es devolverlos en el órden que se declararon
func saludar5(nombre, apellido string, edad uint8) (string, bool) {
	mensaje := fmt.Sprintf("Hola %s %s, tienes %d años", nombre, apellido, edad)
	mayor := edad > 17

	return mensaje, mayor
}

//Se pueden colocar también sólo return, automáticamente
//retorna en el órden indicado
//Pero es necesario declarar las var en la declaración
// de los retornos
func saludar6(nombre, apellido string, edad uint8) (mensaje string, mayor bool) {
	mensaje = fmt.Sprintf("Hola %s %s, tienes %d años", nombre, apellido, edad)
	mayor = edad > 17

	return
}

func saludar7(nombre, apellido string, edad uint8) (mensaje string, mayor bool, err error) {
	mensaje = fmt.Sprintf("Hola %s %s, tienes %d años", nombre, apellido, edad)
	mayor = edad > 17

	if len(apellido) <= 2 {
		err = fmt.Errorf("El apellido es muy corto")
	}

	return
}

//Para recibir varios nombres lo hacemos de la
//siguiente manera
func slice(nombres ...string) {
	fmt.Println(reflect.TypeOf(nombres))
	for _, val := range nombres {
		fmt.Println(val)
	}
}
