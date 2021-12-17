package main

import "fmt"

func main() {
	forma := "rectangulo"

	areaF := getAreaFunc(forma)

	if areaF == nil {
		fmt.Println("forma no soportada")
		return
	}

	res := areaF(2.0, 4.0)
	fmt.Println(res)
}

func getAreaFunc( shape string) func(float64, float64) float64 {

	switch shape {
	case "rectangulo":
		return func(base, altura float64) float64 {
			return base * altura
		}
	case "triangulo":
		return func(base, altura float64) float64 {
			return base * altura / 2
		}
	default:
		return nil
	}
}