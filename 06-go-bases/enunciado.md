Programación Go
# Manejo de errores en Go
Go Bases

---

### Objetivo
El objetivo de esta guía práctica es que podamos afianzar los conceptos sobre el manejo de errores y pkg errors, vistos en el módulo de Go Bases. Para esto planteamos una serie de ejercicios simples e incrementales (ya que vamos a trabajar y agregar complejidad a lo que tenemos que construir), lo que nos permitirá repasar los temas que estudiamos. 

### Forma de trabajo
Para resolver los ejercicios, tengan en cuenta que deben ser realizados en sus computadoras. Les recordamos que generen una carpeta para cada clase y, ahí dentro, tengan un archivo .go para cada ejercicio.


¿Are you ready? 😎👍

---

### 💡 Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el mensaje “Debe pagar impuesto”.

### 💡 Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000. La validación debe ser hecha con la función “Is()” dentro del “main”.


### 💡 Ejercicio 3 - Impuestos de salario #3
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.

### 💡 Ejercicio 4 - Impuestos de salario #4
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible, el mensaje mostrado por consola deberá decir: 
```
Error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]
```
Siendo `[salary]` el valor de tipo int pasado por parámetro.


### 💡 Ejercicio 5 -  Impuestos de salario #5
Vamos a hacer que nuestro programa sea un poco más complejo y útil. 
1. Desarrollá las funciones necesarias para permitir a la empresa calcular:
    1. Salario mensual de un trabajador según la cantidad de horas trabajadas.
        * La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
        * Dicha función deberá retornar más de un valor (salario calculado y error).
        * En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
        * En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo tendrá que indicar “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
