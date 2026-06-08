package data

import "fmt"

func Slice() {
	sliceTest_2()
}

func sliceTest_1() {
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice)
	subSlice := mySlice[1:3]
	fmt.Println(subSlice)
	fmt.Println(mySlice)
	//หาจำนวนทีใช้
	fmt.Println(len(subSlice))
	//หาจำนวนที่มีทั้งหมด
	fmt.Println(cap(subSlice))
}

func sliceTest_2() {
	mySlice := []int{}
	mySlice = append(mySlice, 10)
	fmt.Println(mySlice)
	mySlice = append(mySlice, 20, 30, 10)
	fmt.Println(mySlice)
}
