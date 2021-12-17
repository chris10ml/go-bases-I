package main

import (
	"fmt"
)

type Persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Direccion string
	FechaNac  Fecha
}

type Fecha struct {
	Dia, Mes, Anio int
}

func main() {
	fmt.Println("Estructuras 2")

	//Primera Forma
	p1 := Persona{"Pablo", "Lopez", 27, "Calle 3", Fecha{18, 4, 1994}}
	fmt.Printf("\nP1: %+v, %T", p1, p1)

	//Segunda Forma
	p2 := Persona{}
	p2.Nombre = "Juan"
	p2.FechaNac = Fecha{18, 4, 1994}
	fmt.Printf("\n\nP2: %+v, %T", p2, p2)

	//Tercera Forma
	p3 := Persona{
		Nombre:    "Mariano",
		Direccion: "Calle 1",
		Apellido:  "Bustamante",
		FechaNac: Fecha{
			Dia:  1,
			Mes:  2,
			Anio: 1990,
		},
	}
	fmt.Printf("\n\nP3: %+v, %T", p3, p3)
	fmt.Printf("\nEl nombre de p3 es: %+v", p3.Nombre)
	fmt.Printf("\nEl dia en el que nacio es: %+v", p3.FechaNac.Dia)

}
