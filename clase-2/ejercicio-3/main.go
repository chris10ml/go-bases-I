package main

import "fmt"

func main() {
	edad := 22
	estaEmpleado := true
	antiguedad := 1
	sueldo := 100000

	if edad > 21  && estaEmpleado && antiguedad > 0{
		if sueldo >= 100000 {
			fmt.Println("Se otorga y no se cobra interés!")
		} else {
			fmt.Println("Se otorga y se cobra un poquitin de interés!")
		}
	} else {
		fmt.Printf("Noo se puede!")
	}
}
