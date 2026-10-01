package studentStruct

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name     string
	number   int
	lectures []string
}

func Create(name string, number int, lectures []string) Student {

	return Student{
		Name:     name,
		number:   number,
		lectures: lectures,
	}
}

func ShowInfo(s Student) {
	fmt.Printf("%+v\n", s) // maraqlıdır!!!
	fmt.Println(s)
	fmt.Println(s.Name)
	fmt.Println(s.number)
	fmt.Println(s.lectures)
}

// receiver function
func (s Student) ShowInfoReceiver() {
	fmt.Printf("%+v\n", s) // maraqlıdır!!!
	fmt.Println(s)
	fmt.Println(s.Name)
	fmt.Println(s.number)
	fmt.Println(s.lectures)
}

func (s *Student) RenameReceiver(newName string) {
	s.Name = newName
}

func CreateWithScan() Student {
	var name string
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)
	fmt.Println(name)

	var numberText string
	fmt.Print("Enter Number: ")
	fmt.Scanln(&numberText)
	fmt.Printf("%s, %T", numberText, numberText)

	number, err := strconv.Atoi(numberText)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%d, %T", number, number)
	}

	var lecturesText string
	fmt.Print("Enter Lectures: ")
	fmt.Scanln(&lecturesText)

	lectures := strings.Split(lecturesText, ",")

	newStudent := Student{
		Name:     name,
		number:   number,
		lectures: lectures,
	}

	return newStudent

}

func SaveToFile(students []Student) {

	file, err := os.Create("students.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	for _, std := range students {
		_, err := file.WriteString(
			fmt.Sprintf(
				"------------------------\n"+
					"Student Name: %s\n"+
					"Student Number: %d\n"+
					"Student Lectures: %v\n"+
					"------------------------\n",
				std.Name, std.number, std.lectures,
			),
		)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Saved to file!!")

}
