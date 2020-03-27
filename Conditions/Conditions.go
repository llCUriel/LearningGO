package main

/*
	== igual a
	!= diferente de
	<  menor que
	>  mayor que
	>= mayor o igual que
	&& and
	|| OR
*/
import (
	"fmt"
)

func main() {
	var numberA, numberB, numberC, m int
	numberA = 10
	numberB = 12
	numberC = 14

	if numberA > numberB && numberA > numberC {
		m = numberA
	} else if numberB > numberA && numberB > numberC {
		m = numberB
	} else if numberC > numberA && numberC > numberB {
		m = numberC
	}

	fmt.Println(m)

}
