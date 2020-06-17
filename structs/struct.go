package main

import "fmt"

type User struct {
	edad           int
	name, lastName string
}

func main() {
	var elmer User
	fmt.Println(elmer)

	fmt.Println(User{name: "Elmer", lastName: "Alay"})

	edgar := User{edad: 26, name: "Edgardo", lastName: "Yupe"}
	fmt.Println(edgar)

	//Debe ir en el orden que se declararon en el struct
	//todos los campos deben ir
	yo := User{20, "yo", "el"}
	fmt.Println(yo)

	//Devuelve un puntero
	el := new(User)
	fmt.Println(el)
	//para acceder a un campo no es necesario utilizar *
	el.name = "Orale"
	fmt.Println(el.name)
}
