package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		minimo   = "minimo"
		promedio = "promedio"
		maximo   = "maximo"
	)

	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(err)

	fmt.Printf("Min: %.1f\n", valorMinimo)
	fmt.Printf("Max: %.1f\n", valorMaximo)
	fmt.Printf("Promedio: %.1f\n", valorPromedio)

}

func minFunc(values ...int) float64 {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return float64(min)
}

func maxFunc(values ...int) float64 {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return float64(max)
}

func promFunc(values ...int) float64 {
	var acc int
	var promedio float64
	for _, value := range values {
		acc = acc + value
		promedio = float64(acc / (len(values)))
	}
	return promedio
}

func operacion(value string) (func(...int) float64, error) {
	switch value {
	case "minimo":
		return minFunc, nil

	case "promedio":
		return maxFunc, nil

	case "maximo":
		return promFunc, nil

	default:
		return nil, errors.New("Error")
	}
}
