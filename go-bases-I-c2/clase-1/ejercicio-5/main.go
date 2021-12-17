package main

import "fmt"

func main() {
	const (
		perro     = "perro"
		gato      = "gato"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	animalPerro, err := Animal(perro)
	fmt.Println(err, "perro")
	animalGato, err := Animal(gato)
	fmt.Println(err, "gato")
	animalHamster, err := Animal(hamster)
	fmt.Println(err, "hamster")

	var cantidad float64
	cantidad += animalPerro(5)
	cantidad += animalGato(8)
	cantidad += animalHamster(2)

	fmt.Printf("Cantidad: %.2f\n", cantidad)
}

func animalPerro(cantidad int) float64 {
	return float64(cantidad * 10)
}

func animalGato(cantidad int) float64 {
	return float64(cantidad * 5)
}

func animalHamster(cantidad int) float64 {
	return float64(cantidad) * .250
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case "perro":
		return animalPerro, nil
	case "gato":
		return animalGato, nil
	case "hamster":
		return animalHamster, nil
	default:
		return nil, fmt.Errorf("animal")
	}
}
