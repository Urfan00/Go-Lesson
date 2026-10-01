package main

import "fmt"

func changeName(m *string) {
	*m = "UrFi"
}


func main() {
	name := "Fidan"
	fmt.Println(name)

	p := &name
	fmt.Println("name addressi: ", &name)
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)

	changeName(p)

	fmt.Println(name)
}