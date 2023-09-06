Programación Go
# Estructuras e Interfaces en Go
Práctica clase 2 - Go Bases

### Objetivo
El objetivo de esta guía práctica es que podamos afianzar los conceptos sobre Estructuras e Interfaces, vistos en el módulo de Go Bases. Para esto vamos a plantear una serie de ejercicios simples e incrementales (ya que vamos a ir trabajando y agregando complejidad a lo que tenemos que construir), lo que nos permitirá repasar los temas que estudiamos. 

### Forma de trabajo
Los ejercicios deben ser realizados en sus computadoras. Les recordamos que generen una carpeta para cada clase y ahí dentro tengan un archivo .go para cada ejercicio.

¿Are you ready? 

---

### Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:
```
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
``````
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.

Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

### Ejercicio 2 - Productos
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.

La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
- Pequeño: solo tiene el costo del producto
- Mediano: el precio del producto + un 3% de mantenerlo en la tienda
- Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

> El porcentaje de mantenerlo en la tienda es en base al precio del producto.

> El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

> Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

> Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga
