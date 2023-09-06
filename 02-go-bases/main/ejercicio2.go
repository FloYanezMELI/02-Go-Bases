package main

import "fmt"

func prestamo(edad int, is_empleado bool, antiguedad_trabajo float64, sueldo int){
	if edad > 22 && is_empleado == true && antiguedad_trabajo > 1 && sueldo > 100000{
		fmt.Println("Usted ha sido otorgado un préstamo sin interés")
	} else if edad > 22 && is_empleado == true && antiguedad_trabajo > 1 {
		fmt.Println("Usted ha sido otorgado un préstamo con interés")
	} else {
		fmt.Println("Usted no ha sido otorgado un préstamo")
	}
}