package data

import "fmt"

func changeValue() {
	x := 10

	var p *int = &x

	fmt.Println("value x", x)
	fmt.Println("value p", *p)

	*p = 30

	fmt.Println("value x", x)
	fmt.Println("value p", *p)
}

func Show1() {
	changeValue()
}

type Employee struct {
	Name   string
	Salary int
}

// Function to give a raise to an employee
func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

func Show3() {
	emp := Employee{Name: "John Doe", Salary: 50000}

	giveRaise(&emp, 5000)
	fmt.Println("After raise:", emp)
}
