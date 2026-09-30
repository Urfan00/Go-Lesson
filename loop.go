package main

import "fmt"


func main(){

	// x := 0
	// for x < 5 {
	// 	fmt.Println("x'in degeri: ", x)
	// 	x++
	// }

	// for i := 1; i < 10; i++ {
	// 	fmt.Println("i'nin degeri: ", i)
	// }
	
	students := []string{
		"Urfan",
		"Ayxan",
		"Samil",
	}

	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

	for index, value := range students {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	for _, value := range students {
		fmt.Println(value)

	}

}
