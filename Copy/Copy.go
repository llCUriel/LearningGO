package main

import "fmt"

func main() {
	/*
		Copia el minimo de elementos
	*/
	slice := []int{1, 2, 3, 4, 5, 6}
	_copy := make([]int, len(slice), cap(slice)*2)

	copy(_copy, slice)

	fmt.Println(slice)
	fmt.Println(_copy)

}
