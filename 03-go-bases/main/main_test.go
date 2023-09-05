// https://playground.digitalhouse.com/course/1f38dad9-3a8d-4a1b-b6c3-818f36310664/unit/36c20045-2c4f-49af-93c5-d17185eff5aa/lesson/6d1bd15c-4248-40ce-8032-d71f7eea2a9e/topic/d6fe1295-e4b0-48ab-893b-af07fb361aa7
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Ejercicio 1
func TestImpuestos(t *testing. T){
	// Arrange
	impuesto1 := 43200.0
	impuesto2 := 10200.0
	impuesto3 := 0.0

	// Act
	resultado1 := impuestos(160000)
	resultado2 := impuestos(60000)
	resultado3 := impuestos(40000)
	
	// Assert
	assert.Equal(t, impuesto1, resultado1, "Deben ser iguales")
	assert.Equal(t, impuesto2, resultado2, "Deben ser iguales")
	assert.Equal(t, impuesto3, resultado3, "Deben ser iguales")
}

// Ejercicio 2
func TestPromedio(t *testing. T){
	// Arrange
	promedio1 := 7.0
	promedio2 := 1.0
	promedio3 := 3.5

	// Act
	resultado1 := promedio(7,7,7,7,7)
	resultado2 := promedio(1,1,1,1,1,1,1,1)
	resultado3 := promedio(0,0,0,7,7,7)
	
	// Assert
	assert.Equal(t, promedio1, resultado1, "Deben ser iguales")
	assert.Equal(t, promedio2, resultado2, "Deben ser iguales")
	assert.Equal(t, promedio3, resultado3, "Deben ser iguales")
}

// Ejercicio 3
func TestSalario(t *testing. T){
	// Arrange
	resultadoEsperado1 := 4500.0
	resultadoEsperado2 := 1800.0
	resultadoEsperado3 := 1000.0

	// Act
	resultado1 := salario(60, "A")
	resultado2 := salario(60, "B")
	resultado3 := salario(60, "C")
	
	// Assert
	assert.Equal(t, resultadoEsperado1, resultado1, "Deben ser iguales")
	assert.Equal(t, resultadoEsperado2, resultado2, "Deben ser iguales")
	assert.Equal(t, resultadoEsperado3, resultado3, "Deben ser iguales")
}

// Ejercicio 4
func TestOperation(t *testing. T){
	// Arrange
	resultadoEsperado1 := 0.0
	resultadoEsperado2 := 3.5
	resultadoEsperado3 := 7.0

	// Act
	op1, _ := operation(minimum)
	op2, _ := operation(average)
	op3, _ := operation(maximum)

	resultado1 := op1(6.5,7,1,0,1,6)
	resultado2 := op2(0,0,0,7,7,7)
	resultado3 := op3(6.5,7,1,0,1,6)
	
	// Assert
	assert.Equal(t, resultadoEsperado1, resultado1, "Deben ser iguales")
	assert.Equal(t, resultadoEsperado2, resultado2, "Deben ser iguales")
	assert.Equal(t, resultadoEsperado3, resultado3, "Deben ser iguales")
}

// Ejercicio 5
func TestFoodDog(t *testing. T){
	// Arrange
	resultadoEsperado := 100.0

	// Act
	op, _ := animal(dog)

	resultado := op(10.0)
	
	// Assert
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}

func TestFoodCat(t *testing. T){
	// Arrange
	resultadoEsperado := 50.0

	// Act
	op, _ := animal(cat)

	resultado := op(10.0)
	
	// Assert
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}

func TestFoodHamster(t *testing. T){
	// Arrange
	resultadoEsperado := 2.5

	// Act
	op, _ := animal(hamster)

	resultado := op(10.0)
	
	// Assert
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}

func TestFoodTarantula(t *testing. T){
	// Arrange
	resultadoEsperado := 1.25

	// Act
	op, _ := animal(tarantula)

	resultado := op(10.0)
	
	// Assert
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}