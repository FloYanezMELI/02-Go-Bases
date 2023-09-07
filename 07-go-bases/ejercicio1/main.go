package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando... ")

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err) // Se imprime PORQUE hay un panic
		}
		fmt.Println("Ejecución finalizada")
	}()

	_, err := os.Open("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
 }
 