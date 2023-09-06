package main

func impuestos(sueldo float64) float64{
	if sueldo > 150000 {
		return sueldo * 0.27
	} else if sueldo > 50000 {
		return sueldo * 0.17
	} else {
		return 0
	}
}