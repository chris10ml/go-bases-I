/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar
productos y retornar el valor del precio total.

Las empresas tienen 3 tipos de productos:
- Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda,
y costos de envío.

Sus costos adicionales son:
- Pequeño: El costo del producto (sin costo adicional)
- Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén
de la tienda.
- Grande: El costo del producto + un 6% por mantenimiento, y un costo adicional por
envío de $2500.

Requerimientos:
- Crear una estructura “tienda” que guarde una lista de productos.
- Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
- Crear una interface “Producto” que tenga el método “CalcularCosto”
- Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
- Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y
precio y devuelva un Producto.
- Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
- Interface Producto:
- El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
- Interface Ecommerce:
- El método “Total” debe retornar el precio total en base al costo total de los
productos y los adicionales si los hubiera.
- El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

*/

package main

import "fmt"

// --- Estructuras ---
type tienda struct {
	productos []Producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

// --- Interfaces ---
type Producto interface {
	calcularCosto() float64
}

type Ecommerce interface {
	total() float64
	agregar() float64
}

// --- Main ----
func main() {
	p1 := nuevoProducto("Tecnología", "Mouse Genius", 1200.99)
	fmt.Println(p1)
}

// --- Metodos ---
func nuevoProducto(tipo, nombre string, precio float64) producto {
	return producto{tipo, nombre, precio}
}

func nuevaTienda(productos []Producto) tienda {
	return tienda{productos}
}
