package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Elmer Alay")
	fmt.Println("Que aburrido")

	var wait string
	fmt.Scanln(&wait)
}

func miNombreLento(name string) {
	//La cadena vacía significa que va a separar
	//la cadena en todos sus caracteres
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
