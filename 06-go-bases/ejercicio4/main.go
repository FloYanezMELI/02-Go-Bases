package main

import (
	"fmt"
	"errors"
)

var errorSalario error = errors.New("Error: el mínimo imponible es de 150.000 y el salario ingresado es de:")

func main(){
	fmt.Println("Ejercicio 4:")
	
	salary1 := 5000
	salary2 := 100000

	if salary1 <= 10000 {
		err := fmt.Errorf("%w %d", errorSalario, salary1)
		if errors.Is(err, errorSalario) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Debe pagar")
	}

	if salary2 <= 10000 {
		err := fmt.Errorf("%w %d", errorSalario, salary2)
		if errors.Is(err, errorSalario) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Debe pagar")
	}

}
