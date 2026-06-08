package chapter2

import "fmt"

func Test_if() {
	test_if_3()
}

func test_if_1() {
	var score int = 90

	if score >= 70 {
		fmt.Println("PASSD")
	} else {
		fmt.Println("FAILED")
	}
}

func test_if_2() {
	var score int = 50
	var gread string
	if score >= 80 {
		gread = "A"
	} else if score >= 70 {
		gread = "B"
	} else if score >= 60 {
		gread = "C"
	} else {
		gread = "F"
	}
	fmt.Println("You Aread =", gread)
}

func test_if_3() {
	var num1 int = 5
	var num2 int = 10
	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("> 10")
	}
}

func test_switch_1() {
	var dayOfWeek int = 9
	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday ")
	case 3:
		fmt.Println("Wednesday ")
	case 4:
		fmt.Println("Thursday ")
	case 5:
		fmt.Println("Friday ")
	case 6:
		fmt.Println("Saturday ")
	case 7:
		fmt.Println("Sunday ")
	default:
		fmt.Println("NaN")
	}
}

func test_switch_2() {
	var dayOfWeek int = 3
	var day string
	switch dayOfWeek {
	case 1:
		dayOfWeek = 1
		day = "Monday"
	case 2:
		dayOfWeek = 1
		day = "Tuesday"
	case 3:
		dayOfWeek = 1
		day = "Wednesday"
	case 4:
		dayOfWeek = 1
		day = "Thursday"
	case 5:
		dayOfWeek = 1
		day = "Friday"
	case 6:
		dayOfWeek = 1
		day = "Saturday"
	case 7:
		dayOfWeek = 1
		day = "Sunday "
	default:
		fmt.Println("NaN")
	}
	fmt.Println("dayOfWeek =", day)
}
