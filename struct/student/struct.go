package studentStruct

import "fmt"

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