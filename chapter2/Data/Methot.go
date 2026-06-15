package data

import (
	"fmt"
)

type Student struct {
	Firstname string
	Lastname  string
}

func (s Student) FullName() string {
	return s.Firstname + s.Lastname
}

func Show() {
	student := Student{
		Firstname: "Mike",
		Lastname:  " Lopster",
	}

	// Call the FullName method on the Student instance
	fullName := student.FullName()
	fmt.Println("Full Name of the student:", fullName)
}

type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

func AddData() Rectangle {
	rect := Rectangle{Length: 10, Width: 5}
	return rect
}

func ShowdataMethot() {
	Show := AddData()
	fmt.Println(Show)
	fmt.Println(Show.Area())
}
