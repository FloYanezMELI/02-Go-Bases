package main

import (
	"errors"
)

const (
	dog    		= "dog"
	cat    		= "cat"
	hamster 	= "hamster"
	tarantula 	= "tarantula"
)

func foodDog(cantidad float64) float64 {
	return cantidad * 10
}

func foodCat(cantidad float64) float64 {
	return cantidad * 5
}

func foodHamster(cantidad float64) float64 {
	return cantidad * 0.25
}

func foodTarantula(cantidad float64) float64 {
	return cantidad * 0.125
}

func animal(animal string) (func(cantidad float64) float64, error) {
	switch animal{
	case dog:
		return foodDog, nil
	case cat:
		return foodCat, nil
	case hamster:
		return foodHamster, nil
	case tarantula:
		return foodTarantula, nil
	default:
		return nil, errors.New("Animal inválido")
	}
}


