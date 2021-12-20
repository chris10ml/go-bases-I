/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una
estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:

- Set: Recibe una serie de valores de punto flotante e inicializa los valores en la
	estructura Matrix
- Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos
	de línea entre filas)

La estructura Matrix debe contener los valores de la matriz, la dimensión del alto,
la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main

import "fmt"

type Matrix struct {
	values       []float64
	alto         int
	ancho        int
	esCuadratica bool
	maximo       int
}

func (m Matrix) Set(values []float64) {
	m.values = values
}

func main() {
	m1 := Matrix{
		values:       make([]float64, 2, 5),
		alto:         2,
		ancho:        2,
		esCuadratica: true,
		maximo:       3,
	}

	fmt.Println(m1)

}
