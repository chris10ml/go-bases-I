package main

import "fmt"

func main() {
	//Forma 1
	var p1 *int
	fmt.Printf("p1: %p\n", p1)

	//Forma 2
	var p2 = new(int)
	fmt.Printf("p2: %p\n", p2)

	//Forma 3
	variable := 10
	p3 := &variable
	fmt.Printf("p3: %p\n", p3)

	// Funciones
	Incrementar(&variable)

	fmt.Printf("Incremento a: %d\n", variable)
}

func Incrementar(v *int) {
	*v++
}
