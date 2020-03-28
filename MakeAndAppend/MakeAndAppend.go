package main

import "fmt"

func main() {
	slice := make([]int, 7, 7)
	slice = append(slice, 2)
	fmt.Println(slice)

}
