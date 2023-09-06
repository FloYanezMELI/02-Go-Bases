package main

import "fmt"

func promedio(notas ...int) float64 {
	var sum float64
	var n float64 = float64(len(notas))
	for _, nota := range notas {
		sum += float64(nota)
	}
	var promedio = sum / n
	fmt.Println("Notas:", notas, "\nPromedio:", promedio)
	return promedio
}