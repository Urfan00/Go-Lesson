package main

import "fmt"

func main() {


	// Array
	students := [3]string{
		"Urfan",
		"Ayxan",
		"Samil",
	}

	fmt.Println(students)
	fmt.Println(students[1])
	fmt.Println(len(students))

	students[0] = "Sadiq"
	fmt.Println(students)

	// Slice
	sliceStudent := []string{
		"Urfan",
		"Ayxan",
		"Samil",
	}
	fmt.Println(sliceStudent)
	fmt.Println(sliceStudent[1])
	fmt.Println(len(sliceStudent))

	sliceStudent = append(sliceStudent, "Orxan", "Hesen", "Ali")
	fmt.Println(sliceStudent)
	fmt.Println(len(sliceStudent))


	fmt.Println(sliceStudent[1:4])

	sliceStudent = sliceStudent[0:3]
	fmt.Println(sliceStudent)

}
