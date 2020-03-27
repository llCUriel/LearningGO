package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 23
	strage := strconv.Itoa(age)
	age, _ = strconv.Atoi(strage)
	fmt.Println(age + 30)
	fmt.Println("My age is " + strage)
}
