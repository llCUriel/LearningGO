package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Bla bla bla bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (this Tutor) hablar() string {
	return this.Human.hablar() + " Welcome"
}

func main() {
	tutor := Tutor{Human{"Uriel"}, Dummy{"Uriel"}}
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.hablar())
}
