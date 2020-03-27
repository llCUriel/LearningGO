package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := "Cesar"
	fmt.Printf("Hello %v \n", str)
	var yourAge int
	fmt.Println("What's your age?")
	fmt.Scanf("%d\n", &yourAge)
	fmt.Printf("Your age is %d\n", yourAge)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What's your favorite animal?")
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(text)
	}

}
