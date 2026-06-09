package data

import "fmt"

type student struct {
	Name   string
	Height float32
	Weight float32
	Grade  string
}

func StructTest() {
	structTest_4()

}

func structTest_1() {

	var addstudent student
	addstudent.Name = "Sitthipak"
	addstudent.Weight = 60
	addstudent.Height = 180
	addstudent.Grade = "F"

	fmt.Println(addstudent)
	fmt.Println(addstudent.Grade)
}

func structTest_2() {
	var addstudent [3]student
	addstudent[0].Name = "Sitthipak"
	addstudent[0].Weight = 60
	addstudent[0].Height = 180
	addstudent[0].Grade = "F"

	fmt.Println(addstudent)
	fmt.Println(addstudent[0].Grade)
}

func structTest_3() {

	var addstudent [3]student
	addstudent[0].Name = "Sitthipak"
	addstudent[0].Weight = 60
	addstudent[0].Height = 180
	addstudent[0].Grade = "F"

	for i := 0; i < len(addstudent); i++ {
		fmt.Println("student :", i+1)
		fmt.Println(addstudent[i].Name)
	}
}

func structTest_4() {
	students := make(map[string]student)

	// Add Student structs to the map
	students["st01"] = student{Name: "Mikelopster", Weight: 60, Height: 180, Grade: "F"}
	students["st02"] = student{Name: "Alice", Weight: 55, Height: 165, Grade: "A"}
	students["st03"] = student{Name: "Bob", Weight: 68, Height: 175, Grade: "B"}

	// Print the map
	fmt.Println("Students Map:", students)

	// Access and print a specific student by key
	fmt.Println("Student st01:", students["st01"])
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

// Another struct type used in Person
type Address struct {
	Street  string
	City    string
	ZipCode int
}
