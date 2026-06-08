package data

import "fmt"

func Data() {
	dataTest_4()
}

func dataTest_1() {
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	fmt.Println("Apples:", myMap["apple"])
	fmt.Println("Apples:", myMap)
}

func dataTest_2() {
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8
	//ใช้ Loop ดึงข้อมูลออกมา
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

func dataTest_3() {
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8
	//ลบข้อมูลที่ระบบออกจาก map
	delete(myMap, "orange")
	//ใช้ Loop ดึงข้อมูลออกมา
	for key, value := range myMap {
		fmt.Println(key, value)
	}

}

func dataTest_4() {
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8
	val, ok := myMap["banana"]
	if ok {
		fmt.Println("banana value:", val)
	} else {
		fmt.Println("banana not found in map")
	}

}
