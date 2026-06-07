package chapter2

import (
	"fmt"
)

func Varible() {
	varible_5()
}

func varible_1() {
	//แบบที่ 1
	// var fullName string
	// var age int
	fmt.Println("Hello Word")
}

func varible_2() {
	//แบบที่ 2
	var fullName string = "Tang"
	var age int = 24
	fmt.Println("Hello Word", fullName, "Age", age)
}

func varible_3() {
	//แบบที่ 3
	var fullName = "Tang"
	var age = 25
	fmt.Println("Hello Word", fullName, "Age", age)
}

func varible_4() {
	//แบบที่ 4
	fullName := "Tang"
	age := 25
	fmt.Println("Hello Word", fullName, "Age", age)
	fullName = "Sitthipak"
	fmt.Println("Hello Word", fullName)
}

func varible_5() {
	//แบบที่ 4
	const fullName = "sitthipak"
	const age = 25
	fmt.Println("Hello Word", fullName, "Age", age)
	fmt.Println("Hello Word", fullName)
}
