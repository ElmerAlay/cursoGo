package main

import "fmt"

func main() {
	//chan string debe ser igual en la go routine
	channel := make(chan string)

	//channel es el identificador del canal, puede ser cualquiera
	//chan es el tipo, va a ser un canal
	//String lo que va a enviar
	//chan string debe ser igual a lo que definimos arriba
	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			//Envío informacion
			//channel está del lado izquierdo
			channel <- name
		}
	}(channel)

	//Recibo informacion
	//Channel está del lado derecho
	msg := <-channel
	fmt.Println("Recibido " + msg)
	msg = <-channel
	fmt.Println("Recibido " + msg)
}
