package main

import (
	"fmt"
)

func main() {
	var yourAge int
	fmt.Println("What's your age?")
	fmt.Scanf("%d\n", &yourAge)
	fmt.Printf("%d\n", yourAge)
	Myage := 22
	str := "Cesar"
	flag := true
	fmt.Printf("%v \n", flag)
	fmt.Printf("%v \n", Myage)
	fmt.Printf("%v \n", str)
}
