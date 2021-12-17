package main

import "fmt"

func main() {
	categoria := "B"
	minutosTrabajados := 600

	salario := getSalario(categoria, minutosTrabajados)

	fmt.Printf("El salario final es: $%.2f\n", salario)
}

func getSalario(categoria string, minutos int) (salario float64) {
	horasTrabajadas := float64(minutos / 60)

	switch categoria {
	case "C":
		salario = horasTrabajadas * 1000
	case "B":
		salario = horasTrabajadas * 1500
		salario = salario * 1.2
	case "A":
		salario = horasTrabajadas * 3000
		salario = salario * 1.5
	default:
		salario = 0.0
	}

	return salario
}