package main

import "fmt"

func main() {

	menu := map[string]int{
		"pizza":  20,
		"burger": 15,
		"fries":  5,
	}

	fmt.Println(menu["pizza"])

	fmt.Println(menu["burger"])
	menu["burger"] = 17
	fmt.Println(menu["burger"])

	phoneBook := map[int]string{
		515574569: "Urfan",
		555574569: "Ayxan",
		505574569: "Samil",
	}
	fmt.Println(phoneBook[515574569])

	for k, v := range phoneBook {
		fmt.Printf("%d : %s\n", k, v)
	}

}
