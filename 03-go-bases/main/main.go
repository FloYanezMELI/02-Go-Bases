package main

import (
	"fmt"
)

func main() {
	// Ejemplo
	fmt.Println("Ejemplo")
	resultado := suma(4, 8)
	fmt.Println("El resultado es:", resultado)

	// Ejercicio 1
	fmt.Println("\nEjercicio 1")

	var sueldo1 float64= 160000
	fmt.Println("Sueldo:", sueldo1, "\nImpuestos:", impuestos((sueldo1)))

	var sueldo2 float64 = 60000
	fmt.Println("Sueldo:", sueldo2, "\nImpuestos:", impuestos(sueldo2))

	var sueldo3 float64 = 10000
	fmt.Println("Sueldo:", sueldo3, "\nImpuestos:", impuestos((sueldo3)))


	// Ejercicio 2
	fmt.Println("\nEjercicio 2")
	promedio(1,1,7,7)
	promedio(3,5,6,7)


	// Ejercicio 3
	fmt.Println("\nEjercicio 3")
	salario(2400, "A")
	salario(3000, "B")
	salario(2000, "C")
	salario(1000, "D")

	// Ejercicio 4
	fmt.Println("\nEjercicio 4")
	someFunc, err := operation("fridi")
	if someFunc == nil {
		fmt.Println(err)
	} else {
		someValue := someFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println(someValue)
	}

	minFunc, err := operation(minimum)
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println(minValue)

	averageFunc, err := operation(average)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(averageValue)
	
	maxFunc, err := operation(maximum)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(maxValue)

	// Ejercicio 5
	fmt.Println("\nEjercicio 5")
	var amount float64

	animalDog, msg := animal(dog)
	amount += animalDog(10)

	animalCat, msg := animal(cat)
	amount += animalCat(10)

	animalHamster, msg := animal(hamster)
	amount += animalHamster(10) 

	animalTarantula, msg := animal(tarantula)
	amount += animalTarantula(10)

	fmt.Println(amount)
	fmt.Println(msg)

}

// Ejemplo
func suma(a, b int) int {

	return a + b
}