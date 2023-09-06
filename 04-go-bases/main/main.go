package main

import "fmt"



func main() {
	// Ejercicio 1
	fmt.Println("Ejercicio 1:")
	producto1 := Product{5, "Colet", 1500.0, "Colet de tela con diseño floreado rojo marca Isadora", "Accesorios"}
	producto1.Save()
	producto1.GetAll()
	fmt.Println(getById(2))

	// Ejercicio 2
	fmt.Println("\nEjercicio 2:")
	person := Person{0, "Flo", "09/08/1996"}
	employee1 := Employee{person, 1, "Software Developer"}
	employee1.PrintEmployee()
	employee2 := Employee{Person{1, "Eric", "23/02/2001"}, 2, "Software Developer"}
	employee2.PrintEmployee()
}
