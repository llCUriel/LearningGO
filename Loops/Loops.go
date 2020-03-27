package main

import "fmt"

func main() {
	//1
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			//fmt.Printf("%d x %d = %d \n", i, j, i*j)
		}
	}
	//2
	i := 1
	j := 1
	for i <= 10 {
		j = 1
		for j <= 10 {
			//fmt.Printf("%d x %d = %d \n", i, j, i*j)
			j++
		}
		i++
	}
	//3
	k := 1
	o := 1
	for {
		o = 1
		for {
			//fmt.Printf("%d x %d = %d \n", k, o, k*o)
			o++
			if o > 10 {
				break
			}
		}
		k++
		if k > 10 {
			break
		}
	}
	//4
	for i := 0; i <= 10; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
