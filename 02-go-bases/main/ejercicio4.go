package main

import "fmt"

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