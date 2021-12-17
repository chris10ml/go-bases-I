package main

import "fmt"

func main() {
	sueldo := 150000

	res := getImpuestoSalario(sueldo)

	fmt.Println("El descuento final es: $", res)
}

func getImpuestoSalario(sueldo int) int {
	var descuentoUno int
	var descuentoDos int

	if sueldo >= 50000 {
		descuentoUno = sueldo * 17 / 100
		if sueldo >= 150000 {
			descuentoDos = sueldo * 10 / 100
		}
	}

	return descuentoUno + descuentoDos
}