/*
Una importante empresa de ventas web necesita agregar una funcionalidad
para agregar productos a los usuarios. Para ello requieren que tanto
los usuarios como los productos tengan la misma dirección de memoria
en el main del programa como en las funciones.

Se necesitan las estructuras:
	- Usuario: Nombre, Apellido, Correo, Productos (array de productos).
	- Producto: Nombre, precio, cantidad.
Se requieren las funciones:
	- Nuevo producto: recibe nombre y precio, y retorna un producto.
	- Agregar producto: recibe usuario, producto y cantidad, no retorna nada,
	agrega el producto al usuario.
	- Borrar productos: recibe un usuario, borra los productos del usuario.
*/

package main

import "fmt"

// --- Estructuras ---
type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

// --- Funciones ---
func nuevoProducto(nombre string, precio float64) (p Producto) {
	p.Nombre = nombre
	p.Precio = precio

	return p
}

func agregarProducto(u *Usuario, p *Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, *p)
}

// func borrarProducto(u *Usuario) {
	
// }


func main() {
	p1 := nuevoProducto("Mouse", 900.99)
	fmt.Printf("%+v\n", p1)

	u1 := Usuario{
		Nombre:   "Chris",
		Apellido: "Silvero",
		Correo:   "chris@mail.com",
	}

	agregarProducto(&u1, &p1, 10)

	fmt.Printf("%+v\n", u1)
}
