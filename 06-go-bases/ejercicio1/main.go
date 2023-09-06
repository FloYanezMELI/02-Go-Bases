package main

import "fmt"

type miError1 struct {
	message string
}

func (e miError1) Error() string {
	return e.message
}

func impuesto(sueldo int) string {
	if sueldo <150000 {
		return miError1{"Error: el salario ingresado no alcanza el mínimo imponible"}.Error()
	} else {
		return "Debe pagar"
	}
}

func main(){
	fmt.Println("Ejercicio 1:")
	
	salary1 := 100000
	salary2 := 1000000
	salary3 := 5000

	fmt.Println(impuesto(salary1))
	fmt.Println(impuesto(salary2))
	fmt.Println(impuesto(salary3))
}