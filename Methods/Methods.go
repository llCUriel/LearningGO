package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (this User) nombre_completo() string {
	return this.nombre + " " + this.apellido
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

	fmt.Println("probando")
	fmt.Println(cesar.nombre_completo())
}
