package main

import "fmt"

func salario(minutos float64, categoria string) float64{
	var horas float64 = minutos/60
	var salario float64

	fmt.Println("\nMinutos trabajados:", minutos, "\nHoras trabajadas:", horas, "\nCategoría:", categoria)
	
	switch categoria {
	case "A":
		salario = (horas * 3000) * 1.5
	case "B":
		salario = (horas * 1500) * 1.2
	case "C":
		salario = horas * 1000
	default:
		fmt.Println("Categoría inválida")
	}

	fmt.Println("Salario:", salario)
	return salario
}