package main

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande = "grande"
)

type Pequeño struct {
	precioBase float64
}

func (p Pequeño) precio() float64 {
	return p.precioBase
}

type Mediano struct {
	precioBase float64
}

func (p Mediano) precio() float64 {
	return p.precioBase * 1.03
}

type Grande struct {
	precioBase float64
}

func (p Grande) precio() float64 {
	return p.precioBase * 1.06 + 2500
}

type Producto interface {
	precio() float64
}

func factory(tipo string, precioBase float64) Producto {
	switch tipo {
	case pequeño:
		return Pequeño{precioBase}
	case mediano:
		return Mediano{precioBase}
	case grande:
		return Grande{precioBase}
	}
	return nil
}