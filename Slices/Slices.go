package main

import "fmt"

func main() {
	matrix := []int{1, 2, 3, 4}
	var ma []int
	if matrix == nil {
		fmt.Println("empty")
	} else {
		fmt.Println(matrix)
	}

	if ma == nil {
		fmt.Println("empty")
	} else {
		fmt.Println(ma)
	}
	fmt.Println(len(matrix))

	array := [3]int{1, 2, 3}
	slice := array[1:3]
	fmt.Println(slice)
}
