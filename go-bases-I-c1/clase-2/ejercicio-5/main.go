package main

import "fmt"

func main() {
	listadoAlumnos := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	numbers := make([]int, 6)

	listadoAlumnos = append(listadoAlumnos, "Christian")
	fmt.Println(listadoAlumnos)
	fmt.Printf("Capacidad: %d\n", cap(listadoAlumnos))
	fmt.Printf("Tamaño: %d\n", len(listadoAlumnos))

	fmt.Printf("Numeros: %d", numbers)
}
