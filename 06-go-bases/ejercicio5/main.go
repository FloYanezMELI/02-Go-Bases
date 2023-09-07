package main

import (
	"fmt"
	"errors"
)

var errorHoras error = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")

func salario(horasTrabajadas float64, valorHora float64) (salarioCalculado float64, err error) {
	if horasTrabajadas < 80 {
		return 0.0, errorHoras
	}

	salarioCalculado = horasTrabajadas * valorHora
	if salarioCalculado >= 150000 {
		salarioCalculado = salarioCalculado * 0.9
	}
	return salarioCalculado, nil
}


func main(){
	fmt.Println("Ejercicio 5:")
	
	// Error
	horas := 45.0
	valor := 10000.0

	salarioMensual, err := salario(horas, valor)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salarioMensual)
	}

	// Sin impuesto
	horas = 100.0
	valor = 1400.0

	salarioMensual, err = salario(horas, valor)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salarioMensual)
	}

	// Con impuesto
	horas = 100.0
	valor = 1500.0

	salarioMensual, err = salario(horas, valor)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salarioMensual)
	}

}
