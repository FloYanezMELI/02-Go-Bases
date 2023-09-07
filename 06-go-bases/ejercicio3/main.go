package main

import (
	"fmt"
	"errors"
)

var errorSalario error = errors.New("Error: el salario es menor a 10.000")

func main(){
	fmt.Println("Ejercicio 3:")

	salary1 := 5000
	salary2 := 100000

	if salary1 <= 10000 {
		err := errorSalario
		if errors.Is(err, errorSalario) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Debe pagar")
	}

	if salary2 <= 10000 {
		err := errors.New("Error: el salario es menor a 10.000")
		if errors.Is(err, errors.New("Error: el salario es menor a 10.000")) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Debe pagar")
	}

}
