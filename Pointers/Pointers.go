package main

import "fmt"

func main() {
	/*
		1. Es una dirección de memoria
		2. En lugar del vaolr, tenemos la dirección en la ue está en el valor
		3. X,Y =>as123d =>
		4. X => AS123D => 6
		5. Y=>6
		6. *T => *tipoDeDato
		7. Valor zero => nil
	*/

	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)

}
