package chapter2

import "fmt"

func Operator() {
	operator3()
}

func operator() {
	//Go ไม่อนุญาตให้นำตัวเลขคนละชนิดมาคำนวณกันโดยตรง แม้ว่าจะเป็นตัวเลขเหมือนกันก็ตาม
	var number float32 = 12
	var number2 float32 = 11
	var number3 int = 12
	var number4 int = 11

	fmt.Println(number + number2)
	fmt.Println(number - number2)
	fmt.Println(number * number2)
	fmt.Println(number / number2)
	fmt.Println(number3 % number4)
	fmt.Println("****************")
	number = number * 2
	fmt.Println(number)
}

func operator2() {

	var number float32 = 12
	var number2 float32 = 11

	fmt.Println(number == number2)
	fmt.Println(number != number2)
	fmt.Println(number > number2)
	fmt.Println(number < number2)
	fmt.Println(number >= number2)
	fmt.Println(number <= number2)

}

func operator3() {

	var number bool = true
	var number2 bool = false

	fmt.Println(number && number2)
	fmt.Println(number || number2)
	fmt.Println(!number)

}
