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
