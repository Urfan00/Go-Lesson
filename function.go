package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hello", name)
}


func calcArea(r float64) float64 {
	return 3.14 * r * r
}


func main() {
	sayHello("Urfan")
	sayHello("Fidan")

	area1 := calcArea(5)

	fmt.Println("Area is", area1)
}
