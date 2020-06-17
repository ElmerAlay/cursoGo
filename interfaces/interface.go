package main

import "fmt"

type User interface {
	Permission() int
	Name() string
}

type Admin struct {
	name string
}

//Como es interfaz no debe ser con apuntadores
func (a Admin) Permission() int {
	return 5
}

func (a Admin) Name() string {
	return a.name
}

type Editor struct {
	name string
}

//Como es interfaz no debe ser con apuntadores
func (e Editor) Permission() int {
	return 3
}

func (e Editor) Name() string {
	return e.name
}

func auth(u User) string {
	permission := u.Permission()

	if permission > 4 {
		return u.Name() + " tiene permisos de administrador"
	} else if permission == 3 {
		return u.Name() + " tiene permisos de editor"
	}
	return ""
}

func main() {
	admin := Admin{"Elmer"}
	editor := Editor{"Alay"}
	fmt.Println(auth(admin))
	fmt.Println(auth(editor))

	users := []User{admin, editor}
	for _, user := range users {
		fmt.Println(auth(user))
	}
}
