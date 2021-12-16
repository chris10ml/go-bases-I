package main

import (
	"fmt"
)

func main() {
	var precio int = 180
	var descuento int = 10
	var valorDelDescuento int = precio * descuento / 100
	precioFinal := precio - valorDelDescuento

	fmt.Printf("El precio final con el descuento del %d es: $%d \n", descuento, precioFinal)
}
