package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (first User) nombrecompleto() string {
	return first.nombre + " " + first.apellido
}

func (second User) setname1(name string) {
	second.nombre = name
}

func (third *User) setname2(name string) {
	third.nombre = name
}

func main() {
	wUser := new(User) //pointer
	uriel := User{edad: 23, nombre: "Uriel", apellido: "Hernandez"}
	cesar := User{23, "Cesar", "Castellanos"}
	fmt.Println(wUser)
	fmt.Println(cesar)
	fmt.Println(uriel)

	wUser.edad = 22
	(*wUser).edad = 22
	if wUser.edad == (*wUser).edad {
		fmt.Println("y")
	}

	fmt.Println(cesar.nombrecompleto())

	cesar.setname1("Marcos")
	fmt.Println(cesar.nombre)
	cesar.setname2("Marcos")
	fmt.Println(cesar.nombre)
}
