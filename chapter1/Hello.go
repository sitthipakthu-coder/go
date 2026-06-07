package chapter1

import (
	"fmt"

	"github.com/google/uuid"

	testsayhello "learngo/chapter1/internal/testSayHello"
)

func Say() {
	hello()
	creatid()
	testsayhello.Testsayhello()
}

func hello() {
	fmt.Println("Hello Word")
}

func creatid() {
	id := uuid.New()
	fmt.Println("Id :", id)
}
