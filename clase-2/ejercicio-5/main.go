package main

import "fmt"

func main() {
	listadoAlumnos := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	listadoAlumnos = append(listadoAlumnos, "Christian")
	fmt.Println(listadoAlumnos)
}
