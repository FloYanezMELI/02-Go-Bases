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

// Ejemplo
func ObtenerEstacion(meses []string) {
	for _, mes := range meses {
		if mes == "Enero" || mes == "Febrero" || mes == "Marzo" {
			fmt.Println("Verano")
		} else if mes == "Abril" || mes == "Mayo" || mes == "Junio" {
			fmt.Println("Otoño")
		} else if mes == "Julio" || mes == "Agosto" || mes == "Septiembre" {
			fmt.Println("Invierno")
		} else if mes == "Octubre" || mes == "Noviembre" || mes == "Diciembre" {
			fmt.Println("Primavera")
		} else {
			fmt.Println("No existe este mes")
		}
	}
}


// Rehacer ejemplo con Switch
func ObtenerEstacionSwitch(meses []string) {
	for _, mes := range meses {
		switch mes{
		case "Enero", "Febrero", "Marzo":
			fmt.Println("Verano")
		case "Abril", "Mayo", "Junio":
			fmt.Println("Otoño")
		case "Julio", "Agosto", "Septiembre":
			fmt.Println("Invierno")
		case "Octubre", "Noviembre", "Diciembre":
			fmt.Println("Primavera")
		default:
			fmt.Println("Mes inválido")
		}
	}
}

// Ejercicio 1
func letras(palabra string){
	fmt.Println(palabra)
	fmt.Println(len(palabra))
	for _, letra := range palabra{
		fmt.Println(string(letra))
	}
}

// Ejercicio 2
func prestamo(edad int, is_empleado bool, antiguedad_trabajo float64, sueldo int){
	if edad > 22 && is_empleado == true && antiguedad_trabajo > 1 && sueldo > 100000{
		fmt.Println("Usted ha sido otorgado un préstamo sin interés")
	} else if edad > 22 && is_empleado == true && antiguedad_trabajo > 1 {
		fmt.Println("Usted ha sido otorgado un préstamo con interés")
	} else {
		fmt.Println("Usted no ha sido otorgado un préstamo")
	}
}

// Ejercicio 3
func mes(mes int){
	fmt.Println(mes)
	switch mes{
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("Mes inválido")
	}
}

// Ejercicio 4
func empleados(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Benjamin es", employees["Benjamin"])

	var count int
	for _, edad := range employees{
		if edad > 21{
			count++
		}
	}
	fmt.Println("Cantidad de empleados mayores de 21:", count)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)

}