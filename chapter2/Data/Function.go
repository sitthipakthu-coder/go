package data

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func SayHello2(name string, round int) {
	for i := 0; i < round; i++ {
		fmt.Println("Hello", name)
	}

}

func Add(a int, b int) int {
	return a + b
}

func Adddata() int {
	number1 := 5
	number2 := 8
	SumNumber := Add(number1, number2)
	return SumNumber

}

func ShowAddData() {
	result := Adddata()
	fmt.Println(result)
}
