/*
GUARDAR ARCHIVO

Una empresa que se encarga de vender productos de limpieza necesita:
	1. Implementar una funcionalidad para guardar un archivo de texto, con la información
	de productos comprados, separados por punto y coma (csv).
	2. Debe tener el id del producto, precio y la cantidad.
	3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

package main

import (
	"os"
)

func main() {
	head := []byte("id,precio,cantidad;\n10,800,10;\n12,1230,589;\n")
	os.WriteFile("./data.csv", head, 0644)
}
