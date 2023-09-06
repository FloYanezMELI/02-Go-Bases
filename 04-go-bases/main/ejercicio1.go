package main

import "fmt"

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
