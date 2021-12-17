package main

import "fmt"

func main() {
	numeroDeMes := 10

	switch numeroDeMes {
	case 1:
		fmt.Println(" Enero	31 días")
	case 2:
		fmt.Println(" Febrero	28 días o 29 en año bisiesto")
	case 3:
		fmt.Println(" Marzo	31 días")
	case 4:
		fmt.Println(" Abril	30 días")
	case 5:
		fmt.Println(" Mayo	31 días")
	case 6:
		fmt.Println(" Junio	30 días")
	case 7:
		fmt.Println(" Julio	31 días")
	case 8:
		fmt.Println(" Agosto	31 días")
	case 9:
		fmt.Println(" Septiembre 30 días")
	case 10:
		fmt.Println(" Octubre 31 días")
	case 11:
		fmt.Println(" Noviembre	30 días")
	case 12:
		fmt.Println(" Diciembre	31 días")
	}
}
