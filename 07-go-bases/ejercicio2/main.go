package main

import (
	"bufio"
	"fmt"
	"log"
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

	file, err := os.Open("07-go-bases/ejercicio2/customers.txt") // Tiene esta direccion pq mi proyecto principal es la raíz del repo
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	defer file.Close()

	// Si es que sí se pudo leer el archivo
    scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
 }
 