package main

import "fmt"

type Person struct {
	ID int
	Name string
	DateOfBirth string
}

type Employee struct {
	Person
	ID int
	Position string
}

func (e Employee) PrintEmployee() {
	fmt.Println("ID de Persona:", e.Person.ID)
	fmt.Println("Nombre:", e.Person.Name)
	fmt.Println("Fecha de Nacimiento:", e.Person.DateOfBirth)
	fmt.Println("ID de Empleado:", e.ID)
	fmt.Println("Posición:", e.Position)
}
