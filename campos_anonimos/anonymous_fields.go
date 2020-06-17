package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) hablar() string {
	return "bla bla bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (t *Tutor) hablar() string {
	return "Hola"
}

func main() {
	tutor := Tutor{Human{"Elmer"}, Dummy{"Edgardo"}}

	//Aunque tutor no tenga un campo name
	//el campo anonimo nos permite obtener su campo name
	//Esto sólo funciona si la struct Tutor no tiene
	//otro campo que se llame igual a name

	//fmt.Println(tutor.name)

	//para resolver esto se hace de la siguiente manera
	fmt.Println(tutor.Human.name)

	//Con los métodos no es igual
	//Si tenemos un método en el hijo
	//y luego llamamos de la misma forma a un método
	//del padre el método que tomará siempre será el del padre
	//a menos que especifiquemos
	//Si en dado caso los 2 hijos tienen un método llamado
	//igual se repite el caso de arriba
	fmt.Println(tutor.hablar())
	fmt.Println(tutor.Human.hablar())
}
