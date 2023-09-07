package main

import (
	"fmt"
	"errors"
)

type errorSalario struct {
	mensaje string
}

func (e errorSalario) Error() string {
	return e.mensaje
}

func main(){
	fmt.Println("Ejercicio 2:")

	salary1 := 5000
	salary2 := 100000

	if salary1 <= 10000 {
		err := errorSalario{"Error: el salario es menor a 10.000"}
		if errors.Is(err, errorSalario{"Error: el salario es menor a 10.000"}) {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Debe pagar")
	}

	if salary2 <= 10000 {
		err := errorSalario{"Error: el salario es menor a 10.000"}
		if errors.Is(err, errorSalario{"Error: el salario es menor a 10.000"}) {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Debe pagar")
	}

}
