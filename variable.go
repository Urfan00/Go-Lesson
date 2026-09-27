package main

import "fmt"

func main() {


	// String variables
	var name string = "Urfan"
	var surname = "Aghazada"
	var fatherName string
	motherName := "Aynur"

	fmt.Println(name, surname, fatherName, motherName)

	// Int variables
	var age int = 27
	var weight = 60
	var footSize int
	var height = 1.65

	fmt.Println(age, weight, footSize, height)

	// Float variables
	var balance float64 = 100.5
	var price = 25.5
	var discount float64
	var tax = 18

	fmt.Println(balance, price, discount, tax)

	// Bool variables
	var isStudent bool = true
	var hasJob = false
	var hasLicense bool
	var isMarried = false

	fmt.Println(isStudent, hasJob, hasLicense, isMarried)


	var a, b int = 10, 20

	fmt.Println(a, b)

	c, d := "Salam", 20

	fmt.Println(c, d)

}

