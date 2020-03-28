package main

import (
	"fmt"
)

func main() {
	var myFirstArray [10]int
	mySecondArray := [3]int{0, 1, 2}
	fmt.Println(mySecondArray)
	fmt.Println(myFirstArray)
	fmt.Println(len(myFirstArray))

	i := 0
	for i < len(myFirstArray) {
		myFirstArray[i] = 4
		i++
	}

	fmt.Println(myFirstArray)

	var myMatrix [10][10]string

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			myMatrix[i][j] = "Love is power"
		}
	}
	fmt.Println(myMatrix)

}
