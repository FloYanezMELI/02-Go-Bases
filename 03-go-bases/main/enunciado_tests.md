Bootcamp Go 
# Unit Test en Go
### Objetivo
El objetivo de esta guía práctica es poder afianzar y profundizar los conceptos sobre test unitarios. Para esto, planteamos una serie de ejercicios simples que permitan repasar los temas estudiados. 
### Forma de trabajo
Luego de haber resuelto la práctica de Funciones en Go podrán comenzar el siguiente ejercicio, realizando los test unitarios correspondientes a cada parte de las consignas.

¿Are you ready?

---

### Ejercicio 1 - Testear el impuesto del salario
La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados, al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos. Para esto nos encargaron el trabajo de realizar los test correspondientes para:
- Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
- Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
- Calcular el impuesto en caso de que el empleado gane por encima de $150.000.

### Ejercicio 2 - Calcular promedio
El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente, por lo que ahora nos corresponde realizar los test correspondientes:
- Calcular el promedio de las notas de los alumnos.

### Ejercicio 3 - Test del salario
La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios, por ello nos piden realizar una serie de tests sobre nuestro programa. Necesitaremos realizar las siguientes pruebas en nuestro código:
- Calcular el salario de la categoría “A”.
- Calcular el salario de la categoría “B”.
- Calcular el salario de la categoría “C”.

### Ejercicio 4 - Testear el cálculo de estadísticas
Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google, el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:
- Realizar test para calcular el mínimo de calificaciones.
- Realizar test para calcular el máximo de calificaciones.
- Realizar test para calcular el promedio de calificaciones.

Ejercicio 5 - Calcular cantidad de alimento
El refugio de animales envió una queja ya que el cálculo total de alimento a comprar no fue el correcto y compraron menos alimento del que necesitaban. Para mantener satisfecho a nuestro cliente deberemos realizar los test necesarios para verificar que todo funcione correctamente:
- Verificar el cálculo de la cantidad de alimento para los perros.
- Verificar el cálculo de la cantidad de alimento para los gatos.
- Verificar el cálculo de la cantidad de alimento para los hamster.
- Verificar el cálculo de la cantidad de alimento para las tarántulas.

Ejemplo:
```GO
func TestDog(t *testing.T)
func TestCat(t *testing.T)
``````