/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables
Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func main() {
	a1 := Alumnos{"Chris", "Silvero", "31444431", "13/12/2021"}

	a1.printDetails()
}

func (a Alumnos) printDetails() {
	fmt.Printf("El alumno %s %s con DNI %s, ingresó el %s \n",
		a.Nombre, a.Apellido, a.DNI, a.Fecha)
}
