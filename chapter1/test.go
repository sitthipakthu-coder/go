package chapter1

import (
	"fmt"

	"github.com/google/uuid"
)

func TestHello() {
	fmt.Println("Hello Word test")
	id := uuid.New()
	fmt.Println("Id :", id)
}
