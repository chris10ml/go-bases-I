package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(getPromedio(10, -8))
}

func getPromedio(notas ...int) (float64, error) {
	var sumaDeNotas int
	var promedio float64

	for _, nota := range notas {
		if nota < 0 {
			error := errors.New("Error")
			return 0.0, error
		}
		sumaDeNotas += nota
	}

	promedio = float64(sumaDeNotas / len(notas))

	fmt.Printf("Suma de notas: %d\n", sumaDeNotas)
	fmt.Printf("El promedio: %f\n", promedio)

	return promedio, nil
}
