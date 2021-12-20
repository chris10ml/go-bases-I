/*
LEER ARCHIVO

La misma empresa necesita leer el archivo almacenado, para ello requiere que:
se imprima por pantalla mostrando los valores tabulados,
con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total float64

	// Creo el archivo e inserto datos
	head := []byte("ID,precio,cantidad;1,10,2;2,15,1;3,1,5;")
	os.WriteFile("./data.csv", head, 0644)

	// Leo el archivo
	res, err := os.ReadFile("./data.csv")
	data := strings.Split(string(res), ";")

	if err != nil {
		fmt.Println("Error en la lectura del archivo!")
	}

	for i := 0; i < len(data)-1; i++ {
		var line = strings.Split(string(data[i]), ",")

		if i != 0 {
			precio, err := strconv.ParseFloat(line[1], 64)
			if err != nil {
				println("Precio no se pudo parsear.")
			}
			cantidad, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				println("Cantidad no se pudo parsear.")
			}

			totalProducto := precio * cantidad
			total += totalProducto
		}

		for i := 0; i < len(line); i++ {
			fmt.Printf("%s\t\t", line[i])
			if i == len(line)-1 {
				fmt.Print("\n")
			}
		}
	}

	fmt.Printf("Total\t\t%.2f\n", total)
}
