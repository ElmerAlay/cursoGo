package main

import "fmt"

type User struct {
	age            int
	name, lastName string
}

//Los métodos sólo se pueden agregar a estructuras del mismo
//paquete

//El parámetro puede tener cualquier nombre
//Se puede usar this para que parezca POO
func (u User) nombreCompleto() string {
	return u.name + " " + u.lastName
}

//El nombre no cambiará si ya se había declarado porque esto
//representa otra estructura
//lo mejor es pasarlo como apuntador
func (u User) setName(name string) {
	u.name = name
}

//Ahora sí cambiará el valor porque lo pasamos como apuntador
func (u *User) setName2(name string) {
	u.name = name
}

func main() {
	yo := new(User)
	yo.name = "Elmer"
	yo.lastName = "Alay"

	fmt.Println(yo.nombreCompleto())
	yo.setName("Edgardo")
	fmt.Println(yo.nombreCompleto())
	yo.setName2("Edgardo")
	fmt.Println(yo.nombreCompleto())
}
