package main

import "fmt"

func main() {
	var palabra string = "Chris"
	len := len(palabra)

	fmt.Printf("El tamaño de la palabra \"%s\" es de %d \n", palabra, len)

	for _, character := range palabra {
		fmt.Println(string(character))
	}
}
