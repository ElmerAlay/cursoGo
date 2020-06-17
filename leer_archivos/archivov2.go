package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	executeReadFile()
	fmt.Println("Nunca me voy a ejecutar a menos que use recover")
}

func executeReadFile() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	//Esta forma es mejor para archivos muy grandes
	archivo, err := os.Open("./hola2.txt")

	//Nos aseguramos que siempre se cierre el archivo
	defer func() {
		archivo.Close()
		//lo que venga aquí siempre se va a ejecutar
		fmt.Println("Siempre me ejecuto")

		//Con recover nos recuperamos del panic mode
		//Ahora ya puede ejecutar el demás código
		//también se puede guardar en una variable

		//recover()
		//Nos muestra por qué ocurrió el error
		r := recover()
		fmt.Println(r)
	}()

	if err != nil {
		//Hace que se paniquee el sistema, entonces la función
		//a excepción del defer no va a retornar nada más
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)

	i := 0
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(i)
		i++
		fmt.Println(string(linea))
	}

	return true
}
