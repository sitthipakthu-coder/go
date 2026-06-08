package chapter2

import "fmt"

func Loop() {
	loopTest_3()
}

func loopTest_1() {
	//ForLoop ทำงานจนกว่าจะเป็นเท็จ
	for i := 1; i <= 10; i++ {
		fmt.Println("number:", i)
	}
}

func loopTest_2() {
	//ForLoop ทำงานจนกว่าจะเป็นเท็จ
	for i := 1; i <= 10; i++ {
		//หยุดที่ 5
		if i > 5 {
			break
		}
		fmt.Println("number:", i)
	}
}

func loopTest_3() {
	var i int = 0
	for i <= 10 {
		fmt.Println("number", i)
		i++
	}
}
