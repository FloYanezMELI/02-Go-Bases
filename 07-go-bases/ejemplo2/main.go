package main

import "fmt"

func isPair(num int) {
	defer func() {
		err := recover()
 
		if err != nil {
			fmt.Println(err) // Se imprime PORQUE hay un panic
		}
 
	}()
 
	if (num % 2) != 0 {
		panic("no es un número par")
	}
 
	fmt.Println(num, " es un número par!")
 }

 func main() {
	num := 3
 
	isPair(num)
 
	fmt.Println("Ejecución completada!")
 
 }
 