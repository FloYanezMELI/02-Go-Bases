package main

import (
	"fmt"
	"errors"
)

var errorSalario error = errors.New("Error: el salario es menor a 10.000")

// No usé struct 😟

func checkSalary(salario int) error {
	if salario <= 10000 {
		return errorSalario
	} else {
		return nil
	}
}

func main(){
	fmt.Println("Ejercicio 2:")

	salary1 := 100000
	salary2 := 1000000
	salary3 := 5000

	err1 := checkSalary(salary1)
	err2 := checkSalary(salary2)
	err3 := checkSalary(salary3)

	if errors.Is(err1, errorSalario) {
		fmt.Println(err1)
	} else {
		fmt.Println("Salario ok")
	}

	if errors.Is(err2, errorSalario) {
		fmt.Println(err2)
	} else {
		fmt.Println("Salario ok")
	}

	if errors.Is(err3, errorSalario) {
		fmt.Println(err3)
	} else {
		fmt.Println("Salario ok")
	}
}