package main

import (
	"errors"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )

func opMinimum(notas ...float64) float64 {
	var min float64 = notas[0]
	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}
	return min
}

func opAverage(notas ...float64) float64 {
	var sum float64
	var n float64 = float64(len(notas))
	for _, nota := range notas {
		sum += nota
	}
	return sum / n
}

func opMaximum(notas ...float64) float64 {
	var max float64 = notas[0]
	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}
	return max
}
 
func operation(operacion string) (func(notas ...float64) float64, error) {
	switch operacion{
	case minimum:
		return opMinimum, nil
	case average:
		return opAverage, nil
	case maximum:
		return opMaximum, nil
	default:
		return nil, errors.New("Operación inválida")
	}
}