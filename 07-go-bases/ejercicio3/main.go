package main

import (
	"errors"
	"fmt"
)

type Cliente struct {
	Legajo int
	Nombre string
	DNI string
	Telefono string
	Domicilio string
}

func imprimirClientes() {
	fmt.Println("👤Clientes:")
	for _, c := range Clientes {
		fmt.Println("\t",c)
	}
}
var Clientes []Cliente = []Cliente{
	{1,"Florencia","19475750-6","+56977777777","Calle 123"},
	{2,"Frida","11222333-4","+56988888888","Otra Calle 456"},
}

func (c Cliente) checkIfExists() (bool, error) {
	for _, cliente := range Clientes {
		if c == cliente {
			return false, errors.New("❌ Error: el cliente ya existe.")
		}
	}
	return true, nil
}

func (c Cliente) validar() (bool, error) {
	if c.Legajo == 0 || c.Nombre == "" || c.DNI == "" || c.Telefono == "" || c.Domicilio == "" {
		return false, errors.New("❌ Error: no pueden haber campos vacíos.")
	}
	return true, nil
}

func (c Cliente) crearUsuario() {
	defer func () {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
		fmt.Println("Fin de la ejecución.")
	} ()

	_, err := c.checkIfExists()

	if err != nil {
		panic(err)
	} else {
		_, err = c.validar()
		if err != nil {
			fmt.Println(err)
		} else {
			Clientes = append(Clientes, c)
			fmt.Println("✅ Usuario creado exitosamente")
		}
	}
}

func main() {
	fmt.Println("Iniciando ejecución...")
	imprimirClientes()

	cliente1 := Cliente{3,"Rola","22333444-5","+56999999999","Calle1 111"}
	cliente1.crearUsuario() // Success
	imprimirClientes()
	cliente1.crearUsuario() // Fail pq ya existe

	cliente2 := Cliente{0,"","22333444-5","+56999999999","Calle1 111"}
	cliente2.crearUsuario() // fail pq tiene un campo vacío

}