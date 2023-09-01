package main

import (
	"fmt";
	"errors"
)

// https://playground.digitalhouse.com/course/1f38dad9-3a8d-4a1b-b6c3-818f36310664/unit/36c20045-2c4f-49af-93c5-d17185eff5aa/lesson/2bb68641-55cd-4331-80b3-cbf8827decf3/topic/2325f5c4-812e-4bf1-a5dd-c92a67d23900

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

// Ejercicio 1
func impuestos(sueldo float64) float64{
	if sueldo > 150000 {
		return sueldo * 0.27
	} else if sueldo > 50000 {
		return sueldo * 0.17
	} else {
		return 0
	}
}

// Ejercicio 2
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

// Ejercicio 3
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

// Ejercicio 4
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )

func opMinimum(notas ...float64) float64 {
	var min float64 = notas[0]
	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}
	return min
}

func opAverage(notas ...float64) float64 {
	var sum float64
	var n float64 = float64(len(notas))
	for _, nota := range notas {
		sum += nota
	}
	return sum / n
}

func opMaximum(notas ...float64) float64 {
	var max float64 = notas[0]
	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}
	return max
}
 
func operation(operacion string) (func(notas ...float64) float64, error) {
	switch operacion{
	case minimum:
		return opMinimum, nil
	case average:
		return opAverage, nil
	case maximum:
		return opMaximum, nil
	default:
		return nil, errors.New("Operación inválida")
	}
}

// Ejercicio 5
const (
	dog    		= "dog"
	cat    		= "cat"
	hamster 	= "hamster"
	tarantula 	= "tarantula"
)

func foodDog(cantidad float64) float64 {
	return cantidad * 10
}

func foodCat(cantidad float64) float64 {
	return cantidad * 5
}

func foodHamster(cantidad float64) float64 {
	return cantidad * 0.25
}

func foodTarantula(cantidad float64) float64 {
	return cantidad * 0.125
}

func animal(animal string) (func(cantidad float64) float64, error) {
	switch animal{
	case dog:
		return foodDog, nil
	case cat:
		return foodCat, nil
	case hamster:
		return foodHamster, nil
	case tarantula:
		return foodTarantula, nil
	default:
		return nil, errors.New("Animal inválido")
	}
}


