package main


import "fmt"


func main() {

	age := 20


	// fmt.Println(age >= 25)
	// fmt.Println(age == 27)
	// fmt.Println(age != 25)
	// fmt.Println(age < 25)
	// fmt.Println(age > 25)


	if age >= 25 {
		fmt.Println("You can drive")
	} else if age >=18 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can't drive")
	}
	
}
