package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tipo de dato")

	var numero1 int

	type entero int
	var numero2 entero

	numero1 = 25
	numero2 = 32

	fmt.Printf("\nNumero 1: %v, %T", numero1, numero1)
	fmt.Printf("\nNumero 2: %v, %T", numero2, numero2)

}
