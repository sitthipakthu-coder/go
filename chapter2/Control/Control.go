package chapter2

import "fmt"

func Control() {
	control_2()
}

func control_1() {
	number := 10
	number2 := 20

	fmt.Println(number >= 10 && number2 <= 20)
}

func control_2() {
	name := "Sitthipak"
	lastName := "Thungsiri"

	fmt.Println(name == "Sitthipak" && lastName == "Thungsiri")
}
