package data

import "fmt"

func TestAray() {
	aray3()
}

func aray() {
	var MyArray [3]int
	MyArray[0] = 10
	MyArray[1] = 20
	MyArray[2] = 30
	fmt.Println(MyArray)
	fmt.Println(MyArray[0])
}

func aray2() {
	var MyArray [3]int
	MyArray[0] = 10
	MyArray[1] = 20
	MyArray[2] = 30
	fmt.Println(MyArray)
	fmt.Println(MyArray[0])
	//แก้ไขข้อมูล
	MyArray[0] = 60
	fmt.Println(MyArray[0])
}

func aray3() {
	var MyArray [3]int
	MyArray[0] = 10
	MyArray[1] = 20
	MyArray[2] = 30
	for i := 0; i < len(MyArray); i++ {
		fmt.Println(MyArray[i])
	}

}

// func aray() {
// ทำไมไ่ด้
// 	var MyArray [10,20,30]
// 	fmt.Println(MyArray)
// }
