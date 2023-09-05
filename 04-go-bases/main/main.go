// https://playground.digitalhouse.com/course/1f38dad9-3a8d-4a1b-b6c3-818f36310664/unit/36c20045-2c4f-49af-93c5-d17185eff5aa/lesson/ca9970fd-8956-46ae-b42e-bbe306028e70/topic/53e0b7f8-0bce-4192-8c1e-412db6e69ae2
package main

import "fmt"


// Ejercicio 1
var Products []Product = []Product{
	{0, "Lapiz Azul", 1600.0, "Lapiz de Gel azúl marca Pentel", "Papelería"},
	{1, "Corrector", 2000.0, "Corrector líquido marca PaperMate", "Papelería"},
	{2, "Goma", 500.0, "Goma de borrar marca Faber Castell", "Papelería"},
	{3, "Regla", 800.0, "Regla de medir plástica de 15cm marca Artel", "Papelería"},
	{4, "Mouse", 10000.0, "Mouse inalámbrica a pilas marca Genius", "Tecnología"},
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, prod := range Products{
		fmt.Println(prod)
	}
}

func getById(id int) Product {
	var product Product = Product{}
	for _, prod := range Products {
		if id == prod.ID {
			fmt.Println(prod.ID)
			product = prod
		}
	}
	return product
}

// Ejercicio 2
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
