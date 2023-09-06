package main

import "fmt"

func main() {
	// Ejercicio 1
	fmt.Println("Ejercicio 1:")
	estudiante := Estudiante{"Florencia", "Yáñez", "19.475.750-6", "22/08/2023"}
	estudiante.detalles()

	// Ejercicio 2
	fmt.Println("\nEjercicio 2:")
	fmt.Println(factory(pequeño, 1000.0).precio())
	fmt.Println(factory(mediano, 1000.0).precio())
	fmt.Println(factory(grande, 1000.0).precio())
}