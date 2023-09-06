package main

import "fmt"

func letras(palabra string){
	fmt.Println(palabra)
	fmt.Println(len(palabra))
	for _, letra := range palabra{
		fmt.Println(string(letra))
	}
}