/*
RED SOCIAL

Una empresa de redes sociales requiere implementar una estructura usuario con funciones
que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren
que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
para las funciones:

La estructura debe tener los campos:
Nombre, Apellido, edad, correo y contraseña Y deben implementarse las funciones:
	- cambiar nombre: me permite cambiar el nombre y apellido.
	- cambiar edad: me permite cambiar la edad.
	- cambiar correo: me permite cambiar el correo.
	- cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import "fmt"

type usuario struct {
	nombre      string
	apellido    string
	edad        int
	correo      string
	contrasenia string
}

func (u *usuario) cambiarNombre(nombre, apellido string) {
	u.nombre = nombre
	u.apellido = apellido
}

func (u *usuario) cambiarEdad(edad int) {
	u.edad = edad
}

func (u *usuario) cambiarCorreo(correo string) {
	u.correo = correo
}

func (u *usuario) cambiarContrasenia(contrasenia string) {
	u.contrasenia = contrasenia
}

func main() {
	u1 := usuario{
		nombre:      "Chris",
		apellido:    "Silvero",
		edad:        36,
		correo:      "chris@mail.com",
		contrasenia: "abc1234",
	}

	fmt.Println("Usuario creado:")
	fmt.Println(u1)
	fmt.Println()

	u1.cambiarNombre("Lionel", "Messi")
	u1.cambiarEdad(34)
	u1.cambiarCorreo("lio@messi.com")
	u1.cambiarContrasenia("lio2021")

	fmt.Println("Usuario editado:")
	fmt.Println(u1)
}
