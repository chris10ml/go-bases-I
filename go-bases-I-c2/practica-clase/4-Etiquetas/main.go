package main

import (
	"fmt"
	"reflect"
)

type Fecha struct {
	Dia, Mes, Anio int
}

type Persona struct {
	Nombre    string `json:"user_nombre"`
	Apellido  string `json:"user_apellido"`
	Edad      int    `json:"edad"`
	Direccion string
	FechaNac  Fecha
}

func main() {
	fmt.Println("Etiquetas")

	/*p3 := Persona{
		Nombre:   "Mariano",
		Edad:     25,
		Apellido: "Bustamante",
	}
	fmt.Printf("\n\nP3: %+v, %T", p3, p3)
	fmt.Printf("\nEl nombre de p3 es: %+v", p3.Nombre)

	miJson, err := json.Marshal(p3)
	//fmt.Println("\n", miJson)
	fmt.Println(string(miJson))
	fmt.Println(err)*/

	p5 := Persona{
		Nombre:    "Mariano",
		Direccion: "Calle 1",
		Apellido:  "Bustamante",
		FechaNac: Fecha{
			Dia:  2,
			Mes:  5,
			Anio: 1995,
		},
	}
	fmt.Printf("\n\nP5: %+v, %T", p5, p5)

	fmt.Printf("\n\nReflect\n")
	tipo := reflect.TypeOf(p5)
	valor := reflect.ValueOf(p5)

	for i := 0; i < tipo.NumField(); i++ {
		field := tipo.Field(i)
		fmt.Printf("\n\nMiembro: %+v", field)
		fmt.Printf("\nValor: %+v", valor.Field(i))

		tag := field.Tag.Get("json")
		fmt.Println("\n", tag)
	}

	fmt.Println("\n\nNombre del tipo:", tipo.Name())
	fmt.Println("Es de tipo:", tipo.Kind()) /**/

}
