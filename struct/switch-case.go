package main

import (
	"fmt"
	studentStruct "go-lesson/struct/student"
)

func main() {

	// var text string
	// fmt.Scanln(&text)

	// switch text {
	// case "Salam":
	// 	fmt.Println("Və aleykum salam")
	// case "Hello":
	// 	fmt.Println("Welcome")
	// case "Hola":
	// 	fmt.Println("Bienvenido")
	// default:
	// 	fmt.Println("OKAYY i dont get it")
	// }

	var action string

	data := []studentStruct.Student{}

	fmt.Println("Student Control system")

main_loop:
	for {
		fmt.Println("Please Choose Action: ")
		fmt.Println("1. Add Student")
		fmt.Println("2. List Student")
		fmt.Println("3. Save Student")
		fmt.Println("4. Quit")

		fmt.Scanln(&action)

		switch action {

		case "1":
			new := studentStruct.CreateWithScan()
			data = append(data, new)
		case "2":
			for _, value := range data {
				value.ShowInfoReceiver()
			}
		case "3":
			studentStruct.SaveToFile(data)
		case "4":
			break main_loop
		default:
			fmt.Println("Yanlis secim")
		}
	}

}
