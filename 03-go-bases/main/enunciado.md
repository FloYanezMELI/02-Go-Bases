Programación Go
# Funciones en Go
Práctica clase 2 - Go Bases

---

### Objetivo
El objetivo de esta guía práctica es poder afianzar los conceptos sobre funciones, vistos en el módulo de Go Bases. Para esto vamos a plantear una serie de ejercicios simples e incrementales (trabajaremos y agregaremos complejidad a lo que tenemos que construir), lo que nos permitirá repasar los temas que estudiamos. 

Forma de trabajo
Para resolver los ejercicios, los mismos deben ser realizados en sus computadoras. Les recordamos que generen una carpeta para cada clase y ahí dentro tengan un archivo .go para cada ejercicio.


¿Are you ready? 

---

### Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).


### Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.

### Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
1. Categoría C, su salario es de $1.000 por hora.
2. Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
3. Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.

### Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
Ejemplo:
```GO
const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)

...

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)
 
...

minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
```

### Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:
1. Perro 10 kg de alimento.
2. Gato 5 kg de alimento.
3. Hamster 250 g de alimento.
4. Tarántula 150 g de alimento.

Se solicita:
1. Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
2. Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

Ejemplo:
```GO
const (
   dog    = "dog"
   cat    = "cat"
)

...

animalDog, msg := animal(dog)
animalCat, msg := animal(cat)

...

var amount float64
amount += animalDog(10)
amount += animalCat(10)
```
