package main

import "fmt"


func main(){
	name := "Urfan"
	age := 27
	isSingle := true
	var float float64 = 213234.4353

	// %v -> verb
	fmt.Printf("Hello. My name is %v and I am %v years old.\n", name, age)

	// %q -> dınaqə işarələriylə göstərir
	fmt.Printf("Hello. My name is %q and I am %q years old.\n", name, age)

	// %T -> Tiplerin adini gostermek ucun
	fmt.Printf("Deyiskenin adi: %T\n", isSingle)

	// %f -> Float gorsetmek ucun
	fmt.Printf("Bu floatin deyeri : %f\n", float)

	// %.2f -> 2 reqem qeder gorsetmek ucun
	fmt.Printf("Bu floatin deyeri : %.2f\n", float)

	// %e -> Exponent gorsetmek ucun
	fmt.Printf("Bu floatin deyeri : %e\n", float)

	var myText string = fmt.Sprintf("Hello. My name is %v and I am %v years old.\n", name, age)
	fmt.Println(myText)
}
