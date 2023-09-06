package main

import "fmt"

type Estudiante struct {
	nombre string
	apellido string
	dni string
	fecha string
}

func (e Estudiante) detalles() {
	fmt.Println("Nombre:", e.nombre)
	fmt.Println("Apellido:", e.apellido)
	fmt.Println("DNI:", e.dni)
	fmt.Println("Fecha:", e.fecha)
}