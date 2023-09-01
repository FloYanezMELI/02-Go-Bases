package main

import (
	"fmt"
)

func main() {
	// Ejercicio 1
	var nombre string = "Florencia"
	var direccion string = "Vasco de Gama 4520"
	fmt.Println("Nombre:", nombre, "\nDirección:", direccion)

	// Ejercicio 2
	var temperatura float64 = 20.0
	var humedad float64 = 10.0
	var presion float64 = 9.0
	fmt.Println("Temperatura:", temperatura)
	fmt.Println("Humedad:", humedad)
	fmt.Println("Presión:", presion)

	// Ejercicio 3
	// var 1nombre string				: mal, no puede partir con número
  	// var apellido string				: bien
  	// var int edad						: mal, va el nombre antes del tipo 
  	// 1apellido := 6					: mal, no puede partir con número
  	// var licencia_de_conducir = true	: bien
  	// var estatura de la persona int	: mal, el nombre no puede contener espacios

	// Ejercicio 4
	// var apellido string = "Gomez"	: bien
  	// var edad int = "35"			: mal, debería ser 35 sin comillas
  	// boolean := "false";			: mal, debería ser false sin comillas
  	// var sueldo string = 45857.90	: mal, debería ser float64 en vez de string
  	// var nombre string = "Julián"	: bien

}
