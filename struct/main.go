package main

import (
	"fmt"
	"go-lesson/struct/student"
)

func main() {
	s1 := studentStruct.Create("Fidan", 101, []string{"Math", "Physics"})
	s2 := studentStruct.Create("Urfan", 102, []string{"Math", "History"})

	studentStruct.ShowInfo(s1)

	fmt.Println("*********")
	studentStruct.ShowInfo(s2)

	fmt.Println("*********")
	fmt.Println(s1.Name)

}
