package main

import (
	"fmt"
)

func main() {
	// Ejemplo
	fmt.Println("Ejemplo")
	meses := []string{"Enero", "Febrero", "Junio", "Agosto", "Abril"}
	ObtenerEstacion(meses)
	fmt.Println("Ejemplo con Switch")
	ObtenerEstacionSwitch(meses)

	// Ejercicio 1
	fmt.Println("\nEjercicio 1")
	letras("piña") // tiene 5 letras??
	letras("frida")
	letras("jirafa")
	
	// Ejercicio 2
	fmt.Println("\nEjercicio 2")
	prestamo(21, true, 2, 200000)
	prestamo(40, true, 1.5, 200000)
	prestamo(30, true, 3, 50000)
	
	// Ejercicio 3
	fmt.Println("\nEjercicio 3")
	mes(1)
	mes(2)
	mes(8)
	mes(9)
	mes(10)

	// Ejercicio 4
	fmt.Println("\nEjercicio 4")
	empleados()
}