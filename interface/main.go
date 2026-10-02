package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimetr() float64
}

// duzbucaqli
type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimetr() float64 {
	return 2 * (r.width + r.height)
}

// DAIRE
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimetr() float64 {
	return 2 * math.Pi * c.radius
}

func calculateValues(geo geometry) {
	fmt.Println(geo.area())
	fmt.Println(geo.perimetr())
}

func main() {

	r1 := rect{
		width:  3.5,
		height: 8.4,
	}

	c1 := circle{
		radius: 7,
	}

	// fmt.Println("R area: ", r1.area())
	// fmt.Println("R perimetr: ", r1.perimetr())

	// fmt.Println("C area: ", c1.area())
	// fmt.Println("C perimetr: ", c1.perimetr())

	calculateValues(r1)
	calculateValues(c1)

}
