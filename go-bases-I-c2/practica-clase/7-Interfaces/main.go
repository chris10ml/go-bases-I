package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//detailsCirc(c)

	details(c)
	fmt.Println()
	details(r)

}

//Metodos especificos de Circulo
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*func detailsCirc(c circle) {
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.perim())
}*/

//Metodos especificos de Rectangulo
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func details(g geometry) {
	fmt.Printf("%+T", g)

	fmt.Println("\n", g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
